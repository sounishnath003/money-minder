package migration

import (
	"context"

	"github.com/sounishnath003/money-minder/internal/core"
	"google.golang.org/api/bigquery/v2"
)

// Handles the migration of the services to the database
// google cloud bigquery is used for the migration
// the migration is done in the background
func Migration(co *core.Core) {
	bq, err := bigquery.NewService(context.Background())
	if err != nil {
		co.Logger.Fatalf("Failed to create bigquery client: %v", err)
		return
	}

	dataset, err := bq.Datasets.Insert(co.GoogleProjectId, &bigquery.Dataset{Id: "money_minder_copy", Location: "asia-south1", FriendlyName: "Money Minder Copy", Description: "Money Minder Copy"}).Do()
	if err != nil {
		co.Logger.Fatalf("Failed to create dataset: %v", err)
		return
	}

	co.Logger.Printf("Dataset created: %v", dataset)
}
