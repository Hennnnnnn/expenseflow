package email

import (
	"os"
	"testing"
)

func TestProcessFile(t *testing.T) {

	processor := NewProcessor()

	transaction, err := processor.ProcessFile(
		"../../testdata/bca/credit_card_transaction.eml",
	)

	if err != nil {
		t.Fatal(err)
	}

	if transaction == nil {
		t.Fatal("transaction should not be nil")
	}

	if transaction.Merchant == "" {
		t.Fatal("merchant should not be empty")
	}

	if transaction.Amount <= 0 {
		t.Fatal("amount should be greater than zero")
	}
}

func TestProcessBytes(t *testing.T) {

	processor := NewProcessor()

	data, err := os.ReadFile(
		"../../testdata/bca/credit_card_transaction.eml",
	)
	if err != nil {
		t.Fatal(err)
	}

	transaction, err := processor.ProcessBytes(data)
	if err != nil {
		t.Fatal(err)
	}

	if transaction == nil {
		t.Fatal("transaction should not be nil")
	}

	if transaction.Merchant == "" {
		t.Fatal("merchant should not be empty")
	}

	if transaction.Amount <= 0 {
		t.Fatal("amount should be greater than zero")
	}
}
