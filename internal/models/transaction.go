package models

import "time"

type Transaction struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	TransactionType string `json:"transactionType"`
	CategoryID      int    `json:"categoryID"`
	Amount          int    `json:"amount"`

	UserID    int       `json:"userID"`
	CreatedAt time.Time `json:"createdAt"`
}
