package parser

import "github.com/Hennnnnnn/expenseflow/internal/email/imap"

type Parser interface {
	CanParse(message *imap.Message) bool

	Parse(message *imap.Message) (*TransactionData, error)
}
