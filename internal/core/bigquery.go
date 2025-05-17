package core

import (
	"context"
	"time"

	"cloud.google.com/go/bigquery"
	"github.com/sounishnath003/money-minder/internal/models"
	"google.golang.org/api/iterator"
)

type CustomBigQueryClient struct {
	bigquery.Client
}

func (cbq *CustomBigQueryClient) GetTransactionsByUserId(userId int, fromDate, endDate time.Time) ([]models.Transaction, error) {
	query := cbq.Client.Query(QUERY_GET_TRANSACTIONS_BY_USER_ID)
	query.Parameters = []bigquery.QueryParameter{
		{
			Name:  "UserID",
			Value: userId,
		},
		{
			Name:  "FromDate",
			Value: fromDate,
		},
		{
			Name:  "EndDate",
			Value: endDate,
		},
	}

	iter, err := query.Read(context.Background())
	if err != nil {
		return nil, err
	}

	var transactions []models.Transaction
	for {
		var transaction models.Transaction
		err := iter.Next(&transaction)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}
