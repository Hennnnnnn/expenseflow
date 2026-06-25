package email

import "github.com/Hennnnnnn/expenseflow/internal/email/imap"

type ReaderService interface {
	ReadLatest(limit int) ([]imap.Message, error)
	ReadBody(seq uint32) (*imap.Message, error)
	ReadBCAEmails(limit int) ([]imap.Message, error)
}
