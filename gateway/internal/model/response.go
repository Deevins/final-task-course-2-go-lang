package model

// ErrorResponse описывает ошибку API.
type ErrorResponse struct {
	Error string `json:"error" example:"validation failed"`
}

// TransactionsResponse описывает список транзакций.
type TransactionsResponse struct {
	Transactions []Transaction `json:"transactions"`
}

// BudgetsResponse описывает список бюджетов.
type BudgetsResponse struct {
	Budgets []Budget `json:"budgets"`
}

// ReportsResponse описывает список отчетов.
type ReportsResponse struct {
	Reports []Report `json:"reports"`
}
