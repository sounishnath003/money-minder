package core

import (
	"context"
	"log"

	"cloud.google.com/go/bigquery"
	"github.com/sounishnath003/money-minder/internal/models"
	"github.com/sounishnath003/money-minder/internal/utility"
)

type Core struct {
	Port            int
	Hostname        string
	Logger          *log.Logger
	GoogleProjectId string
	ServerJwtSecret string
	UserStore       map[string]models.User // key: site-public-key

	BQClient *CustomBigQueryClient
}

// InitCore - initializes all the important services and instances requires to run the project
func InitCore() *Core {
	co := &Core{
		Port:            utility.GetNumberValueFromEnv("PORT", 3000),
		Hostname:        utility.GetStringValueFromEnv("HOSTNAME", "0.0.0.0"),
		GoogleProjectId: utility.GetStringValueFromEnv("GOOGLE_PROJECT_ID", "sounish-cloud-workstation"),
		Logger:          log.Default(),

		// JWT Server side site secret key
		ServerJwtSecret: utility.GetJWTSecret(),
	}

	// Initialize the user store context from the Site private secret key
	// Get the site public and private key
	sitePublicKey := utility.GetStringValueFromEnv("SITE_SECRET_PUBLIC_KEY", "default_site_public_secret")
	sitePrivateKey := utility.GetStringValueFromEnv("SITE_SECRET_PRIVATE_KEY", "default_site_private_secret")

	// Update the store context information
	co.UserStore = map[string]models.User{
		sitePublicKey: {
			ID:       1, // 0 is not allowed. 1 is the seed value, do not change
			Name:     "Sounish Nath",
			Password: sitePrivateKey,
		},
	}

	// Initiatize the google bigquery client
	bqClient, err := bigquery.NewClient(context.Background(), co.GoogleProjectId)
	if err != nil {
		log.Panic("not able to initialize the bigquery client: %w", err)
	}
	co.BQClient = &CustomBigQueryClient{*bqClient}

	return co
}
