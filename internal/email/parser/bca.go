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

	return strings.Contains(from, "bca")
}

func (p *BCAParser) Parse(message *imap.Message) (*TransactionData, error) {

	return &TransactionData{}, nil
}
