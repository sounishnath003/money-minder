package models

import "time"

type PaymentMethod string

const (
	Cash        PaymentMethod = "Cash"
	UPI         PaymentMethod = "UPI"
	BankDeposit PaymentMethod = "Bank Deposit"
	CreditCard  PaymentMethod = "Credit Card"
)

type Transaction struct {
	ID              int           `json:"id"`
	Name            string        `json:"name"`
	TransactionType string        `json:"transactionType"`
	PaymentMethod   PaymentMethod `json:"paymentMethod"`
	CategoryID      int           `json:"categoryID"`
	Amount          int           `json:"amount"`

	UserID    int       `json:"userID"`
	CreatedAt time.Time `json:"createdAt"`
}
