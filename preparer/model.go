package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/oklog/ulid/v2"
)

type Record struct {
	Hash    string
	URL     string
	Login   bool
	Created int64
}

func genRecords(count int) []*Record {
	records := make([]*Record, 0, count)
	for i := 0; i < count; i++ {
		records = append(records, &Record{
			Hash:    ulid.Make().String(),
			URL:     fmt.Sprintf("https://example.com/%d", i),
			Created: time.Now().Unix(),
		})
	}

	return records
}

func ensureDynamoDBHashTable(ctx context.Context, dynamoClient *dynamodb.Client) error {
	tableName := aws.String("hash")

	_, err := dynamoClient.DescribeTable(ctx, &dynamodb.DescribeTableInput{
		TableName: tableName,
	})
	if err == nil {
		fmt.Println("dynamodb table TargetedCoupon is already existing")

		return nil
	}

	if temp := new(types.ResourceNotFoundException); !errors.As(err, &temp) {
		return err
	}

	// Table was not created before so do it.
	_, err = dynamoClient.CreateTable(ctx, &dynamodb.CreateTableInput{
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("Hash"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("Hash"),
				KeyType:       types.KeyTypeHash,
			},
		},
		TableName:   tableName,
		BillingMode: types.BillingModePayPerRequest,
	})

	if err != nil {
		return err
	}

	return nil
}
