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

type RawTransactionData struct {
	CustomerNumber  string
	CardNumber      string
	Merchant        string
	TransactionType string
	TransactionDate string
	Amount          string
}
