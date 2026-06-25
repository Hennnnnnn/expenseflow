package service

import (
	"github.com/Hennnnnnn/expenseflow/internal/domain/entity"
	"gorm.io/gorm"
	"errors"
)

type transactionService struct {
	db *gorm.DB
}

func NewTransactionService(db *gorm.DB) TransactionService {
	return &transactionService{
		db: db,
	}
}

func (s *transactionService) Create(tx *entity.Transaction) error {
	
	exists, err := s.Exists(tx.ReferenceNo)

	if err != nil {
		return err
	}

	if exists {
		return errors.New("transaction already exists")
	}
	
	return s.db.Create(tx).Error
}

func (s *transactionService) GetAll() ([]entity.Transaction, error) {

	var transactions []entity.Transaction

	err := s.db.
		Order("transaction_date DESC").
		Find(&transactions).Error

	return transactions, err
}

func (s *transactionService) Exists(referenceNo string) (bool, error) {
	var count int64

	err := s.db.
			Model(&entity.Transaction{}).
			Where("reference_no = ?", referenceNo).
			Count(&count).Error

	return count > 0, err

}