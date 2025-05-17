package models

import "time"

type Transaction struct {
	ID         int       `json:"transactionID"`
	Name       string    `json:"name"`
	CatagoryID string    `json:"transactionCatagory"`
	Timestamp  time.Time `json:"timestamp"`

	UserID int `json:"userID"`
}
