package email

import (
	"github.com/Hennnnnnn/expenseflow/internal/ai/pipeline"
	"github.com/Hennnnnnn/expenseflow/internal/service"
)

type SyncService struct {
	reader ReaderService

	processor *Processor

	transactionService service.TransactionService

	pipeline *pipeline.Service
}

func NewSyncService(
	reader ReaderService,
	processor *Processor,
	transactionSvc service.TransactionService,
	pipeline *pipeline.Service,
) *SyncService {

	return &SyncService{
		reader:             reader,
		processor:          processor,
		transactionService: transactionSvc,
		pipeline:           pipeline,
	}
}

func (s *SyncService) SyncFile(path string) error {

	tx, err := s.processor.ProcessFile(path)
	if err != nil {
		return err
	}

	_, err = s.saveTransaction(tx)
	return err
}

func (s *SyncService) SyncLatest(limit int) (*SyncResult, error) {

	result := &SyncResult{}

	emails, err := s.reader.ReadBCAEmails(limit)
	if err != nil {
		return nil, err
	}

	for _, mail := range emails {
		s.processMail(mail, result)
	}

	return result, nil
}
