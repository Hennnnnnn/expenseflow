package database

import (
	"os"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func NewSQLite() (*gorm.DB, error) {

	if err := os.MkdirAll("data", os.ModePerm); err != nil {
		return nil, err
	}

	return gorm.Open(
		sqlite.Open("data/expenseflow.db"),
		&gorm.Config{},
	)
}