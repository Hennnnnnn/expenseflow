package mapper

import (
	"fmt"
	"time"

	"github.com/Hennnnnnn/expenseflow/internal/domain/entity"
	"github.com/Hennnnnnn/expenseflow/internal/domain/types"
	"github.com/Hennnnnnn/expenseflow/internal/email/parser/bca"
	"github.com/google/uuid"
)

func BCAParserToEntity(data *bca.TransactionData) *entity.Transaction {

	return &entity.Transaction{
		UUID:            uuid.NewString(),
		TransactionDate: data.TransactionDate,
		Amount:          data.Amount,
		Merchant:        data.Merchant,
		Category:        "Uncategorized",

		Type: types.TransactionExpense,

		Source: types.SourceBCAEmail,

		ReferenceNo: GenerateReferenceNo(data),
	}
}

func GenerateReferenceNo(tx *bca.TransactionData) string {
	return fmt.Sprintf(
		"%s|%s|%.0f|%s",
		tx.CardNumber,
		tx.TransactionDate.Format(time.RFC3339),
		tx.Amount,
		tx.Merchant,
	)
}
