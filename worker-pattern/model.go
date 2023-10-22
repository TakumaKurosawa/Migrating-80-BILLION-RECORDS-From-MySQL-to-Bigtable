package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

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
	workerLimit = 100
	channelSize = 200
)

func migrateRecordsForShard(ctx context.Context, db *gorm.DB, dynamoDB *dynamodb.Client, table string, querySize int, shardWG *sync.WaitGroup) {
	defer shardWG.Done()

	insertChan := make(chan *Record, channelSize)
	var workerWg sync.WaitGroup
	for i := 0; i < workerLimit; i++ {
		workerWg.Add(1)
		fmt.Printf("Worker-%d is started!!\n", i)
		go func() {
			if err := insertToDynamo(ctx, dynamoDB, insertChan, &workerWg); err != nil {
				log.Println(err)

				return
			}
		}()
	}

	var lastHash string
	for {
		records, recordLeft, err := getRecordsFromMySQL(db, table, lastHash, querySize)
		if err != nil {
			log.Println(err)

			return
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
	workerWg.Wait()
	fmt.Printf("%s is done!\n", table)
}

func getRecordsFromMySQL(db *gorm.DB, table, lastHash string, limit int) ([]*Record, bool, error) {
	var records []*Record
	if err := db.Table(table).Where("hash > ?", lastHash).Find(&records).Limit(limit).Error; err != nil {
		return nil, false, err
	}

	return records, len(records) < limit, nil
}

func insertToDynamo(ctx context.Context, db *dynamodb.Client, insertChan <-chan *Record, workerWg *sync.WaitGroup) error {
	defer workerWg.Done()

	for record := range insertChan {
		time.Sleep(1 * time.Millisecond) // simulate network latency
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
