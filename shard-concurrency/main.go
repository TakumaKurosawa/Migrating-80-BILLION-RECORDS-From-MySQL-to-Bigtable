package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	shardCount = 6
	querySize  = 20
)

func main() {
	ctx := context.Background()

	db := setupMySQL()
	dynamoDB, err := setupDynamoDB(ctx)
	if err != nil {
		log.Println(err)

		return
	}

	startTime := time.Now()

	fmt.Println("-------- Start migration --------")

	var shardWaitGroup sync.WaitGroup
	for i := 1; i <= shardCount; i++ {
		fmt.Printf("shard-%d is started!!\n", i)

		shardWaitGroup.Add(1)

		go func(i int) {
			var lastHash string
			for {
				var records []*Record
				records, recordLeft, err := getRecordsFromMySQL(db, fmt.Sprintf("hashdb-%d", i), lastHash, querySize)
				if err != nil {
					log.Println(err)

					return
				}

				if err := insertToDynamo(ctx, dynamoDB, records); err != nil {
					log.Println(err)

					return
				}

				lastHash = records[len(records)-1].Hash

				if !recordLeft {
					fmt.Printf("shard-%d is done!\n", i)
					shardWaitGroup.Done()

					break
				}
			}
		}(i)
	}

	shardWaitGroup.Wait()

	fmt.Println("-------- Finish migration --------")
	fmt.Printf("Migration took %s\n", time.Since(startTime))
}

func setupMySQL() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:local-pass@tcp(127.0.0.1:3306)/local?charset=utf8mb4&parseTime=True&loc=UTC"), nil)
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func setupDynamoDB(ctx context.Context) (*dynamodb.Client, error) {
	awsCfg, err := awsconfig.LoadDefaultConfig(
		ctx,
		awsconfig.WithRegion("ap-northeast-1"),
		awsconfig.WithEndpointResolverWithOptions(
			aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{
					URL: "http://localhost:8000",
				}, nil
			}),
		),
		awsconfig.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID:     "dummy",
				SecretAccessKey: "dummy",
				SessionToken:    "dummy",
				Source:          "Hard-coded credentials; values are irrelevant for local DynamoDB",
			},
		}),
	)
	if err != nil {
		return nil, err
	}

	return dynamodb.NewFromConfig(awsCfg), nil
}
