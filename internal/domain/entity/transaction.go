package entity

import "time"

type Transaction struct {
	ID uint `gorm:"primaryKey"`

	TransactionDate time.Time
	Amount          float64

	Type        string
	Merchant    string
	Description string
	Category    string
	Source      string

	ReferenceNo string

	CreatedAt time.Time
	UpdatedAt time.Time
}