package bca

import (
	"fmt"
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

	if value == "" {
		return 0, fmt.Errorf("amount is empty")
	}

	// sementara belum support foreign currency
	if strings.HasPrefix(strings.ToUpper(value), "USD") {
		return 0, fmt.Errorf("foreign currency is not supported")
	}

	value = strings.ReplaceAll(value, "Rp", "")
	value = strings.ReplaceAll(value, ".", "")
	value = strings.ReplaceAll(value, ",00", "")
	value = strings.ReplaceAll(value, " ", "")

	if value == "" {
		return 0, fmt.Errorf("amount is empty after normalization")
	}

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
		CardNumber:      raw.CardNumber,
		Merchant:        raw.Merchant,
		TransactionType: raw.TransactionType,
		TransactionDate: date,
		Amount:          amount,
	}, nil
}
