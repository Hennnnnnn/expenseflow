package service

import (
	"github.com/Hennnnnnn/expenseflow/internal/email"
	"github.com/Hennnnnnn/expenseflow/internal/email/imap"
)

type EmailService struct {
	reader email.Reader
}

func New(reader email.Reader) *EmailService {
	return &EmailService{
		reader: reader,
	}
}

func (s *EmailService) ReadLatest(limit int) ([]imap.Message, error) {
	return s.reader.ReadLatest(limit)
}
