package core

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/bigquery"
	"github.com/sounishnath003/money-minder/internal/models"
	"google.golang.org/api/iterator"
)

type CustomBigQueryClient struct {
	bigquery.Client
}

// GetTransactionsByUserId - helps to get all the transactions between FROM_DATE AND END_DATE by user ID.
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

	if len(transactions) == 0 {
		return []models.Transaction{}, nil
	}

	return transactions, nil
}

// GetActiveCategories helps to get all active category enabled on the service
func (cbq *CustomBigQueryClient) GetActiveCategories() ([]models.Category, error) {
	query := cbq.Query(QUERY_GET_ACTIVE_CATRGORY)

	it, err := query.Read(context.Background())
	if err != nil {
		return nil, err
	}

	var categories []models.Category
	for {
		var category models.Category
		err := it.Next(&category)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

// AddTransaction - helps to add a transaction for a user ID
func (cbq *CustomBigQueryClient) AddTransaction(transaction models.Transaction) (bool, error) {
	// @ID, @Name, @TransactionType, @UserID, @CategoryID, @Amount, @CreatedAt
	query := cbq.Client.Query(QUERY_ADD_NEW_TRANSACTION)
	query.Parameters = []bigquery.QueryParameter{
		{
			Name:  "ID",
			Value: transaction.ID,
		},
		{
			Name:  "Name",
			Value: transaction.Name,
		},
		{
			Name:  "TransactionType",
			Value: transaction.TransactionType,
		},
		{
			Name:  "UserID",
			Value: transaction.UserID,
		},
		{
			Name:  "CategoryID",
			Value: transaction.CategoryID,
		},
		{
			Name:  "Amount",
			Value: transaction.Amount,
		},
		{
			Name:  "CreatedAt",
			Value: transaction.CreatedAt,
		},
	}

	job, err := query.Run(context.Background())
	if err != nil {
		return false, fmt.Errorf("failed to run query: %w", err)
	}

	_, err = job.Wait(context.Background())
	if err != nil {
		return false, fmt.Errorf("failed to wait for job: %w", err)
	}

	return true, nil
}
