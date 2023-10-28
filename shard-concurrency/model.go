package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"gorm.io/gorm"
)

type Record struct {
	Hash    string
	URL     string
	Login   bool
	Created int64
}

func migration(ctx context.Context, db *gorm.DB, dynamoDB *dynamodb.Client, table string) error {
	var lastHash string
	for {
		var records []*Record
		records, recordLeft, err := getRecordsFromMySQL(db, table, lastHash, querySize)
		if err != nil {
			return err
		}

		if err := insertToDynamo(ctx, dynamoDB, records); err != nil {
			return err
		}

		lastHash = records[len(records)-1].Hash

		if !recordLeft {
			break
		}
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

func insertToDynamo(ctx context.Context, db *dynamodb.Client, records []*Record) error {
	for _, record := range records {
		apiRes, err := somethingHeavyTransaction()
		if err != nil {
			return err
		}
		record.URL = apiRes.Image

		//item, err := attributevalue.MarshalMap(record)
		//if err != nil {
		//	return err
		//}
		//
		//if _, err := db.PutItem(ctx, &dynamodb.PutItemInput{
		//	Item:      item,
		//	TableName: aws.String("hash"),
		//}); err != nil {
		//	return err
		//}
	}

	return nil
}

type res struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Image string `json:"image"`
}

func somethingHeavyTransaction() (*res, error) {
	//url := fmt.Sprintf("https://api.sampleapis.com/coffee/hot/%d", rand.Intn(20))
	//resp, err := http.Get(url)
	//if err != nil {
	//	return nil, err
	//}
	//
	//defer resp.Body.Close()
	//
	//var data res
	//if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
	//	return nil, err
	//}

	time.Sleep(100 * time.Millisecond)

	return &res{
		ID:    rand.Intn(20),
		Title: "Response",
		Image: fmt.Sprintf("https://api.sampleapis.com/coffee/hot/%d", rand.Intn(20)),
	}, nil
}
