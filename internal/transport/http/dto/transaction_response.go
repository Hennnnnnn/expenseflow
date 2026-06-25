package dto

type TransactionResponse struct {
	ID              uint    `json:"id"`
	TransactionDate string  `json:"transaction_date"`
	Amount          float64 `json:"amount"`
	Type            string  `json:"type"`
	Merchant        string  `json:"merchant"`
	Description     string  `json:"description"`
	Category        string  `json:"category"`
	Source          string  `json:"source"`
	ReferenceNo     string  `json:"reference_no"`
}