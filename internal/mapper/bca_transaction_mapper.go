package mapper

import (
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

		ReferenceNo: data.ReferenceNumber,
	}
}
