package bca

import "testing"

func TestNormalize(t *testing.T) {

	n := NewNormalizer()

	raw := &RawTransactionData{
		Merchant:        "UKI MATCHA PURI",
		TransactionType: "DOMESTIK",
		Amount:          "Rp201.000,00",
		TransactionDate: "25-06-2026 19:32:45 WIB",
	}

	transaction, err := n.Normalize(raw)

	if err != nil {
		t.Fatal(err)
	}

	if transaction.Amount != 201000 {
		t.Fatalf("expected 201000 got %f", transaction.Amount)
	}

	if transaction.Merchant != "UKI MATCHA PURI" {
		t.Fatal("merchant mismatch")
	}
}
