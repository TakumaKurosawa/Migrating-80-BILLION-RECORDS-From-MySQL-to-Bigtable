package main

import (
	"context"
	"fmt"
	"testing"

	"golang.org/x/sync/errgroup"
)

func BenchmarkMigration(b *testing.B) {
	ctx := context.Background()

	db := setupMySQL()
	dynamoDB, err := setupDynamoDB(ctx)
	if err != nil {
		b.Error(err)

		return
	}

	eg, egCtx := errgroup.WithContext(ctx)

	for i := 1; i <= b.N; i++ {
		table := fmt.Sprintf("hashdb-%d", i)

		eg.Go(func() error {
			if err := migration(egCtx, db, dynamoDB, table, querySize); err != nil {
				return err
			}

			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		b.Error(err)

		return
	}
}
