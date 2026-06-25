package email

import (
	"fmt"

	"github.com/Hennnnnnn/expenseflow/internal/mapper"
	"github.com/Hennnnnnn/expenseflow/internal/service"
)

type SyncService struct {
	reader ReaderService

	processor *Processor

	transactionService service.TransactionService
}

func NewSyncService(
	reader ReaderService,
	processor *Processor,
	transactionSvc service.TransactionService,
) *SyncService {

	return &SyncService{
		reader:             reader,
		processor:          processor,
		transactionService: transactionSvc,
	}
}

func (s *SyncService) SyncFile(path string) error {

	tx, err := s.processor.parser.ParseFile(path)
	if err != nil {
		return err
	}

	entity := mapper.BCAParserToEntity(tx)

	exists, err := s.transactionService.Exists(entity.ReferenceNo)
	if err != nil {
		return err
	}

	if exists {
		return nil
	}

	return s.transactionService.Create(entity)
}

func (s *SyncService) SyncLatest(limit int) error {

	emails, err := s.reader.ReadBCAEmails(limit)
	if err != nil {
		return err
	}

	for _, mail := range emails {

		body, err := s.reader.ReadBody(mail.SeqNum)
		if err != nil {
			return err
		}

		fmt.Println(body.TextBody[:200])

		_ = body
	}

	return nil
}
