package email

import "github.com/Hennnnnnn/expenseflow/internal/email/imap"

type Reader interface {
	Connect() error
	Close() error

	ReadLatest(limit int) ([]imap.Message, error)
	ReadBody(uid uint32) (*imap.Message, error)
}
