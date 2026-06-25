package database

import (
	"github.com/Hennnnnnn/expenseflow/internal/domain/entity"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&entity.Transaction{},
	)
}