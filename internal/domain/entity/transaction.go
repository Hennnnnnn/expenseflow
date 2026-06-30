package entity

import (
	"time"
	"github.com/Hennnnnnn/expenseflow/internal/domain/types"
)

type Transaction struct {
	ID uint `gorm:"primaryKey"`

	UUID string `gorm:"uniqueIndex"`

	TransactionDate time.Time

	Amount float64

	Type types.TransactionType

	Category string

	Merchant string 

	Description string

	ReferenceNo string `gorm:"uniqueIndex"`

	Source types.TransactionSource

	CreatedAt time.Time

	UpdatedAt time.Time
}