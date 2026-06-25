package bca

import "testing"

func TestParseFile(t *testing.T) {

	service := NewService()

	transaction, err := service.ParseFile(
		"../../../../testdata/bca/credit_card_transaction.eml",
	)

	if err != nil {
		t.Fatal(err)
	}

	if transaction == nil {
		t.Fatal("transaction should not be nil")
	}

	t.Logf("%+v", transaction)
}
