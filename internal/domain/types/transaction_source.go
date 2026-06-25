package types

type TransactionSource string

const (
	SourceManual TransactionSource = "MANUAL"
	SourceBCAEmail TransactionSource = "BCA_EMAIL"
	SourceTelegram TransactionSource = "TELEGRAM"
)