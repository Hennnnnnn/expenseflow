package database

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func NewSQLite() (*gorm.DB, error) {
	return gorm.Open(
		sqlite.Open("data/expenseflow.db"),
		&gorm.Config{},
	)
}