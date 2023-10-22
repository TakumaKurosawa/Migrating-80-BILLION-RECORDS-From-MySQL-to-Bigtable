package main

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"gorm.io/gorm"
)

type Record struct {
	Hash    string
	URL     string
	Login   bool
	Created int64
}

func getRecordsFromMySQL(db *gorm.DB, table, lastHash string, limit int) ([]*Record, bool, error) {
	var records []*Record
	if err := db.Table(table).Where("hash > ?", lastHash).Find(&records).Limit(limit).Error; err != nil {
		return nil, false, err
	}

	return records, len(records) < limit, nil
}

func insertToDynamo(ctx context.Context, db *dynamodb.Client, records []*Record) error {
	for _, record := range records {
		item, err := attributevalue.MarshalMap(record)
		if err != nil {
			return err
		}

		if _, err := db.PutItem(ctx, &dynamodb.PutItemInput{
			Item:      item,
			TableName: aws.String("hash"),
		}); err != nil {
			return err
		}

	}

	return nil
}
