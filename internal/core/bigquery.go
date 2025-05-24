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

// SpendByCateogory helps to store the category spend by every item tx.
type SpendByCateogory struct {
	Category    string `json:"category"`
	TotalAmount int    `json:"totalAmount"`
}

// GetTotalSpendsByCategories helps to get the group by aggregation of
// the user's ID fromDate to EndDate total spend by category.
func (cbq *CustomBigQueryClient) GetTotalSpendsByCategories(userId int, fromDate, endDate time.Time) ([]SpendByCateogory, error) {
	query := cbq.Query(QUERY_GET_TRANSACTION_SPEND_BY_CATEGORY)
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

	it, err := query.Read(context.Background())
	if err != nil {
		return nil, err
	}

	var results []SpendByCateogory

	for {
		var spend SpendByCateogory
		err := it.Next(&spend)
		if err == iterator.Done {
			break
		}

		if err != nil {
			return nil, err
		}
		results = append(results, spend)
	}

	return results, nil
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

type DailyTotalSpend struct {
	UnixMiliseconds int `json:"unixMiliseconds"`
	Amount          int `json:"amount"`
}

// GetDailyTotalSpendByTimeframe returns the user's daily spend based
// on the particular from date and end date.
// DailyTotalSpend
func (cbq *CustomBigQueryClient) GetDailyTotalSpendByTimeframe(userId int, fromDate, endDate time.Time) ([]DailyTotalSpend, error) {
	query := cbq.Query(ANALYTICS_TOTAL_DAILY_SPEND_BY_DATE_RANGE)
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

	it, err := query.Read(context.Background())
	if err != nil {
		return nil, err
	}

	var results []DailyTotalSpend

	for {
		var spend DailyTotalSpend
		err := it.Next(&spend)
		if err == iterator.Done {
			break
		}

		if err != nil {
			return nil, err
		}
		results = append(results, spend)
	}

	return results, nil
}

type SpendOnCategoryMoM struct {
	Seq              string `json:"-"`
	Month            string `json:"month"`
	Category         string `json:"category"`
	TotalSpendAmount int    `json:"totalSpendAmount"`
}

// GetSpendOnCategoriesMonthOnMonth helps to get the spend on categories Month on Month from every month
func (cbq *CustomBigQueryClient) GetSpendOnCategoriesMonthOnMonth(userId int, fromDate, endDate time.Time) ([]SpendOnCategoryMoM, error) {
	query := cbq.Query(ANALYTICS_SPENDS_ON_CATEGORIES_MONTH_ON_MONTH)
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

	it, err := query.Read(context.Background())
	if err != nil {
		return nil, err
	}

	var results []SpendOnCategoryMoM

	for {
		var spend SpendOnCategoryMoM
		err := it.Next(&spend)
		if err == iterator.Done {
			break
		}

		if err != nil {
			return nil, err
		}
		results = append(results, spend)
	}

	return results, nil
}
