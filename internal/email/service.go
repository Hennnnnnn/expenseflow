package email

import (
	"github.com/Hennnnnnn/expenseflow/internal/email/imap"
)

type Service struct {
	reader Reader
}

func New(reader Reader) *Service {
	return &Service{
		reader: reader,
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
