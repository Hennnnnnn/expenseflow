package mapper

import (
	"time"

	"github.com/Hennnnnnn/expenseflow/internal/domain/entity"
	"github.com/Hennnnnnn/expenseflow/internal/transport/http/dto"
)

func ToTransactionResponse(tx entity.Transaction) dto.TransactionResponse {
	return dto.TransactionResponse{
		ID:              tx.ID,
		TransactionDate: tx.TransactionDate.Format(time.RFC3339),
		Amount:          tx.Amount,
		Type:            string(tx.Type),
		Merchant:        tx.Merchant,
		Description:     tx.Description,
		Category:        tx.Category,
		Source:          string(tx.Source),
		ReferenceNo:     tx.ReferenceNo,
	}
}

func ToTransactionResponses(transactions []entity.Transaction) []dto.TransactionResponse {
	result := make([]dto.TransactionResponse, 0, len(transactions))

	for _, tx := range transactions {
		result = append(result, ToTransactionResponse(tx))
	}

	return result
}