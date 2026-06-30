package email

import (
	"log"

	"github.com/Hennnnnnn/expenseflow/internal/email/parser/bca"
	"github.com/Hennnnnnn/expenseflow/internal/mapper"
)

func (s *SyncService) saveTransaction(
	tx *bca.TransactionData,
) (bool, error) {

	result := s.pipeline.Categorize(
		tx.Merchant,
	)

	log.Printf(
		"Merchant: %s | Category: %s | Confidence: %.2f",
		tx.Merchant,
		result.Category,
		result.Confidence,
	)

	entity := mapper.BCAParserToEntity(tx)
	entity.Category = result.Category

	exists, err := s.transactionService.Exists(entity.ReferenceNo)
	if err != nil {
		return false, err
	}

	if exists {
		return false, nil
	}

	if err := s.transactionService.Create(entity); err != nil {
		return false, err
	}

	return true, nil
}
