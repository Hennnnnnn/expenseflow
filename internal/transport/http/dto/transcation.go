package dto

type CreateTransactionRequest struct {
	TransactionDate string  `json:"transaction_date" validate:"required"`
	Amount          float64 `json:"amount" validate:"required,gt=0"`
	Type            string  `json:"type" validate:"required"`
	Merchant        string  `json:"merchant"`
	Description     string  `json:"description"`
	Category        string  `json:"category"`
	Source          string  `json:"source"`
	ReferenceNo     string  `json:"reference_no"`
}