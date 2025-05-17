package core

import (
	"context"
	"log"

	"cloud.google.com/go/bigquery"
	"github.com/sounishnath003/money-minder/internal/utility"
)

type Core struct {
	Port            int
	Hostname        string
	Logger          *log.Logger
	GoogleProjectId string

	BQClient *bigquery.Client
}

// InitCore - initializes all the important services and instances requires to run the project
func InitCore() *Core {
	co := &Core{
		Port:            utility.GetNumberValueFromEnv("PORT", 3000),
		Hostname:        utility.GetStringValueFromEnv("HOSTNAME", "localhost"),
		GoogleProjectId: utility.GetStringValueFromEnv("GOOGLE_PROJECT_ID", "sounishnath-cloud-workstation"),
		Logger:          log.Default(),
	}

	// Initiatize the google bigquery client
	bqClient, err := bigquery.NewClient(context.Background(), co.GoogleProjectId)
	if err != nil {
		log.Panic("not able to initialize the bigquery client: %w", err)
	}
	co.BQClient = bqClient

	return co
}
