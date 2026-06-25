package email

import (
	"testing"

	"github.com/Hennnnnnn/expenseflow/internal/database"
	"github.com/Hennnnnnn/expenseflow/internal/service"
)

func TestSyncFile(t *testing.T) {

	db, err := database.NewSQLite()
	if err != nil {
		t.Fatal(err)
	}

	if err := database.AutoMigrate(db); err != nil {
		t.Fatal(err)
	}

	transactionService := service.NewTransactionService(db)

	processor := NewProcessor()

	sync := NewSyncService(
		nil,
		processor,
		transactionService,
	)

	err = sync.SyncFile(
		"../../testdata/bca/credit_card_transaction.eml",
	)

	if err != nil {
		t.Fatal(err)
	}
}
