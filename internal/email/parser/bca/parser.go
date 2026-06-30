package bca

import (
	"errors"
	"os"
	"strings"
)

var ErrNotTransaction = errors.New("not transaction email")

type Parser struct {
}

func New() *Parser {
	return &Parser{}
}

func (p *Parser) FindValue(text string, label string) string {

	lines := strings.Split(text, "\n")

	for i := 0; i < len(lines); i++ {

		current := strings.TrimSpace(lines[i])

		if current != label {
			continue
		}

		for j := i + 1; j < len(lines); j++ {

			value := strings.TrimSpace(lines[j])

			// skip baris kosong
			if value == "" {
				continue
			}

			// skip separator
			if value == ":" {
				continue
			}

			return value
		}
	}

	return ""
}

func (p *Parser) Parse(text string) (*RawTransactionData, error) {

	if !p.CanParse(text) {
		return nil, ErrNotTransaction
	}

	os.WriteFile("debug.txt", []byte(text), 0644)

	println("====================")
	println("Merchant :", p.FindValue(text, "Merchant / ATM"))
	println("Amount   :", p.FindValue(text, "Sejumlah"))
	println("====================")

	return &RawTransactionData{
		CustomerNumber:  p.FindValue(text, "Nomor Customer"),
		CardNumber:      p.FindValue(text, "Nomor Kartu"),
		Merchant:        p.FindValue(text, "Merchant / ATM"),
		TransactionType: p.FindValue(text, "Jenis Transaksi"),
		TransactionDate: p.FindValue(text, "Pada Tanggal"),
		Amount:          p.FindValue(text, "Sejumlah"),
	}, nil
}

func (p *Parser) CanParse(text string) bool {

	return strings.Contains(text, "Merchant / ATM") &&
		strings.Contains(text, "Sejumlah")

}
