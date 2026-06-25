package email

import (
	"github.com/Hennnnnnn/expenseflow/internal/email/filter"
	"github.com/Hennnnnnn/expenseflow/internal/email/imap"
	"github.com/Hennnnnnn/expenseflow/internal/email/parser"
)

type Service struct {
	reader Reader

	parsers []parser.Parser
}

func New(reader Reader) *Service {
	return &Service{
		reader: reader,
		parsers: []parser.Parser{
			parser.NewBCAParser(),
		},
	}
}

func (s *Service) ReadLatest(limit int) ([]imap.Message, error) {
	return s.reader.ReadLatest(limit)
}

func (s *Service) ReadBody(uid uint32) (*imap.Message, error) {
	return s.reader.ReadBody(uid)
}

func (s *Service) PrintLatest(limit int) error {

	messages, err := s.ReadLatest(limit)
	if err != nil {
		return err
	}

	for _, msg := range messages {
		println("======================")
		println(msg.Subject)
		println(msg.From)
		println(msg.Date.String())
	}

	return nil
}

func (s *Service) ReadBCAEmails(limit int) ([]imap.Message, error) {

	messages, err := s.ReadLatest(limit)
	if err != nil {
		return nil, err
	}

	result := make([]imap.Message, 0)

	for _, message := range messages {

		println("--------------------------")
		println("Sender :", message.From)
		println("Subject:", message.Subject)

		if !filter.IsBCA(message.From) {
			println("❌ Sender rejected")
			continue
		}

		if !filter.IsTransaction(message.Subject) {
			println("❌ Subject rejected")
			continue
		}

		println("✅ Accepted")

		result = append(result, message)
	}

	return result, nil
}
