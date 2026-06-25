package email

import "testing"

func TestProcessFile(t *testing.T) {

	processor := NewProcessor()

	err := processor.ProcessFile(
		"../../testdata/bca/credit_card_transaction.eml",
	)

	if err != nil {
		t.Fatal(err)
	}
}
