package main

import (
	"context"
	"log"

	"golang.org/x/sync/errgroup"

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

const (
	workerLimit = 10
	channelSize = 200
)

func migration(ctx context.Context, db *gorm.DB, dynamoDB *dynamodb.Client, table string, querySize int) error {
	insertChan := make(chan *Record, channelSize)
	workerEg, ctx := errgroup.WithContext(ctx)

	for i := 1; i <= workerLimit; i++ {
		workerEg.Go(func() error {
			if err := insertToDynamo(ctx, dynamoDB, insertChan); err != nil {
				return err
			}

			return nil
		})
	}

	var lastHash string
	for {
		records, recordLeft, err := getRecordsFromMySQL(db, table, lastHash, querySize)
		if err != nil {
			log.Println(err)

			return err
		}

		for _, record := range records {
			insertChan <- record
		}

		lastHash = records[len(records)-1].Hash

		if !recordLeft {
			break
		}
	}

	close(insertChan)

	if err := workerEg.Wait(); err != nil {
		return err
	}

	return nil
}

func getRecordsFromMySQL(db *gorm.DB, table, lastHash string, limit int) ([]*Record, bool, error) {
	var records []*Record
	if err := db.Table(table).Where("hash > ?", lastHash).Find(&records).Limit(limit).Error; err != nil {
		return nil, false, err
	}

	return records, len(records) < limit, nil
}

func insertToDynamo(ctx context.Context, db *dynamodb.Client, insertChan <-chan *Record) error {
	for record := range insertChan {
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
