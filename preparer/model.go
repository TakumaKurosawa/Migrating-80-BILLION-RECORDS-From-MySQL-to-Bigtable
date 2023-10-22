package main

import (
	"fmt"
	"github.com/oklog/ulid/v2"
	"time"
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
