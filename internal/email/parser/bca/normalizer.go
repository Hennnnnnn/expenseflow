package bca

import (
	"strconv"
	"strings"
	"time"
)

type Normalizer struct {
}

func NewNormalizer() *Normalizer {
	return &Normalizer{}
}

func (n *Normalizer) NormalizeAmount(value string) (float64, error) {

	value = strings.TrimSpace(value)

	value = strings.ReplaceAll(value, "Rp", "")
	value = strings.ReplaceAll(value, ".", "")
	value = strings.ReplaceAll(value, ",00", "")
	value = strings.ReplaceAll(value, " ", "")

	return strconv.ParseFloat(value, 64)

}

func (n *Normalizer) NormalizeDate(value string) (time.Time, error) {

	layout := "02-01-2006 15:04:05 MST"

	return time.Parse(layout, value)

}

func (n *Normalizer) Normalize(raw *RawTransactionData) (*TransactionData, error) {

	amount, err := n.NormalizeAmount(raw.Amount)
	if err != nil {
		return nil, err
	}

	date, err := n.NormalizeDate(raw.TransactionDate)
	if err != nil {
		return nil, err
	}

	return &TransactionData{
		Merchant:        raw.Merchant,
		TransactionType: raw.TransactionType,
		TransactionDate: date,
		Amount:          amount,
	}, nil
}
