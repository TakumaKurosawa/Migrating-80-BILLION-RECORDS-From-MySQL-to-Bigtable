package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"golang.org/x/sync/errgroup"

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

	eg, ctx := errgroup.WithContext(ctx)

	for i := 1; i <= shardCount; i++ {
		table := fmt.Sprintf("hashdb-%d", i)
		fmt.Printf("%s is starting!!\n", table)

		eg.Go(func() error {
			if err := migration(ctx, db, dynamoDB, table); err != nil {
				return err
			}

			return nil
		})

	}

	if err := eg.Wait(); err != nil {
		log.Println(err)

		return
	}

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
