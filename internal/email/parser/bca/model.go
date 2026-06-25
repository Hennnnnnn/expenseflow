package bca

import "time"

type TransactionData struct {
	Merchant string

	Amount float64

	TransactionDate time.Time

	TransactionType string

	ReferenceNumber string

	CardNumber string

	Description string
}
