package service

import "github.com/Hennnnnnn/expenseflow/internal/domain/entity"

type TransactionService interface {
	Create(tx *entity.Transaction) error
	GetAll() ([]entity.Transaction, error)
	Exists(referenceNo string) (bool, error)
}