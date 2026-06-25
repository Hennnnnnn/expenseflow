package parser

import "time"

type TransactionData struct {
	Amount float64

	Merchant string

	TransactionDate time.Time

	ReferenceNumber string

	Description string

	Category string
}
