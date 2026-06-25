package types

type TransactionType string

const (
	TransactionIncome TransactionType = "INCOME"
	TransactionExpense TransactionType = "EXPENSE"
)