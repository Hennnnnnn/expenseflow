package mapper

import (
	"time"
	"testing"
	"github.com/Hennnnnnn/expenseflow/internal/domain/entity"
	"github.com/Hennnnnnn/expenseflow/internal/domain/types"
)

func TestToTransactionResponse(t *testing.T) {

	tx := entity.Transaction{
		ID:              1,
		TransactionDate: time.Date(2026, 6, 25, 12, 0, 0, 0, time.UTC),
		Amount:          25000,
		Type:            types.TransactionExpense,
		Merchant:        "McDonald's",
		Description:     "Lunch",
		Category:        "Food",
		Source:          types.SourceManual,
		ReferenceNo:     "TEST-001",
	}

	result := ToTransactionResponse(tx)

	if result.ID != tx.ID {
		t.Errorf("expected ID %d, got %d", tx.ID, result.ID)
	}

	if result.Amount != tx.Amount {
		t.Errorf("expected Amount %.2f, got %.2f", tx.Amount, result.Amount)
	}

	if result.Merchant != tx.Merchant {
		t.Errorf("expected Merchant %s, got %s", tx.Merchant, result.Merchant)
	}

	if result.Type != string(tx.Type) {
		t.Errorf("expected Type %s, got %s", tx.Type, result.Type)
	}
}