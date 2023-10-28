package main

import (
	"context"
	"fmt"
	"testing"
)

func BenchmarkMigration(b *testing.B) {
	ctx := context.Background()

	db := setupMySQL()
	dynamoDB, err := setupDynamoDB(ctx)
	if err != nil {
		b.Error(err)

		return
	}

	for i := 1; i <= b.N; i++ {
		if err := migration(ctx, db, dynamoDB, fmt.Sprintf("hashdb-%d", i), querySize); err != nil {
			b.Error(err)
		}
	}
}
