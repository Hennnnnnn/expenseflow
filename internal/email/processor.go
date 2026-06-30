package email

import (
	"github.com/Hennnnnnn/expenseflow/internal/email/parser/bca"
)

type Processor struct {
	parser *bca.Service
}

func NewProcessor() *Processor {
	return &Processor{
		parser: bca.NewService(),
	}
}

func (p *Processor) ProcessFile(path string) (*bca.TransactionData, error) {
	return p.parser.ParseFile(path)
}

func (p *Processor) ProcessBytes(data []byte) (*bca.TransactionData, error) {
	return p.parser.ParseBytes(data)
}
