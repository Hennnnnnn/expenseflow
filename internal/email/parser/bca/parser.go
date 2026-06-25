package bca

import "strings"

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

	return &RawTransactionData{
		CustomerNumber:  p.FindValue(text, "Nomor Customer"),
		CardNumber:      p.FindValue(text, "Nomor Kartu"),
		Merchant:        p.FindValue(text, "Merchant / ATM"),
		TransactionType: p.FindValue(text, "Jenis Transaksi"),
		TransactionDate: p.FindValue(text, "Pada Tanggal"),
		Amount:          p.FindValue(text, "Sejumlah"),
	}, nil
}
