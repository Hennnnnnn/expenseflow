package parser

import (
	"strings"

	"github.com/Hennnnnnn/expenseflow/internal/email/imap"
)

type BCAParser struct {
}

func NewBCAParser() *BCAParser {
	return &BCAParser{}
}

func (p *BCAParser) CanParse(message *imap.Message) bool {

	from := strings.ToLower(message.From)

	subject := strings.ToLower(message.Subject)

	if strings.Contains(from, "bca") {
		return true
	}

	if strings.Contains(subject, "bca") {
		return true
	}

	return false
}

func (p *BCAParser) Parse(message *imap.Message) (*TransactionData, error) {

	return &TransactionData{}, nil
}
