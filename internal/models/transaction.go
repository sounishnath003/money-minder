package models

import "time"

type Transaction struct {
	ID              int    `json:"transactionID"`
	Name            string `json:"name"`
	TransactionType string `json:"transactionType"`
	CategoryID      int    `json:"catagoryID"`

	UserID    int       `json:"userID"`
	CreatedAt time.Time `json:"createdAt"`
}
