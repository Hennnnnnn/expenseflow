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

func (p *Processor) ProcessFile(path string) error {

	transaction, err := p.parser.ParseFile(path)

	if err != nil {
		return err
	}

	println("Merchant :", transaction.Merchant)
	println("Amount   :", transaction.Amount)

	return nil
}
