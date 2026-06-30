package dto

type DashboardResponse struct {
	IncomeThisMonth     float64             `json:"incomeThisMonth"`
	ExpenseThisMonth    float64             `json:"expenseThisMonth"`
	TransactionCount    int                 `json:"transactionCount"`
	TopCategories       []CategorySummary   `json:"topCategories"`
}

type CategorySummary struct {
	Category string  `json:"category"`
	Amount   float64 `json:"amount"`
}