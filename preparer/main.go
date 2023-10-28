package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	shardCount          = 6
	recordCountPerShard = 10000
)

func main() {
	ctx := context.Background()
	db := setupMySQL()
	dynamoDB := setupDynamoDB(ctx)

	fmt.Println("-------- Start create records --------")
	for i := 1; i <= shardCount; i++ {
		fmt.Printf("shard-%d...", i)
		recoords := genRecords(recordCountPerShard)

		if err := db.Table(fmt.Sprintf("hashdb-%d", i)).Create(&recoords).Error; err != nil {
			log.Printf("failed to create records, err: %#v\n", err)

			return
		}
		fmt.Println("done!")
	}
	fmt.Printf("-------- Finish create records --------\n\n")

	fmt.Println("-------- Start dynamodb ensure table --------")
	if err := ensureDynamoDBHashTable(ctx, dynamoDB); err != nil {
		log.Printf("failed to ensure dynamodb table, err: %#v\n", err)

		return
	}
	fmt.Println("Hash table is ready!")
	fmt.Println("-------- Finish dynamodb ensure table --------")
}

func setupMySQL() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:local-pass@tcp(127.0.0.1:3306)/local?charset=utf8mb4&parseTime=True&loc=UTC"), nil)
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func setupDynamoDB(ctx context.Context) *dynamodb.Client {
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
		panic("failed to create dynamodb client")
	}

	return dynamodb.NewFromConfig(awsCfg)
}
