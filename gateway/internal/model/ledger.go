package model

import "time"

// Transaction описывает транзакцию в Ledger.
type Transaction struct {
	ID          string    `json:"id" example:"11111111-1111-1111-1111-111111111111"`
	AccountID   string    `json:"account_id" example:"22222222-2222-2222-2222-222222222222"`
	Amount      float64   `json:"amount" example:"1250.50"`
	Currency    string    `json:"currency" example:"RUB"`
	Category    string    `json:"category" example:"Продукты"`
	Description string    `json:"description" example:"Покупка в магазине"`
	OccurredAt  time.Time `json:"occurred_at" example:"2024-01-01T10:00:00Z"`
	CreatedAt   time.Time `json:"created_at" example:"2024-01-01T10:00:00Z"`
	UpdatedAt   time.Time `json:"updated_at" example:"2024-01-01T10:00:00Z"`
}

// CreateTransactionRequest описывает запрос на создание транзакции.
type CreateTransactionRequest struct {
	AccountID   string    `json:"account_id" example:"22222222-2222-2222-2222-222222222222"`
	Amount      float64   `json:"amount" binding:"required" example:"1250.50"`
	Currency    string    `json:"currency" binding:"required" example:"RUB"`
	Category    string    `json:"category" binding:"required" example:"Продукты"`
	Description string    `json:"description" example:"Покупка в магазине"`
	OccurredAt  time.Time `json:"occurred_at" binding:"required" example:"2024-01-01T10:00:00Z"`
}

// UpdateTransactionRequest описывает запрос на обновление транзакции.
type UpdateTransactionRequest struct {
	AccountID   string    `json:"account_id" example:"22222222-2222-2222-2222-222222222222"`
	Amount      float64   `json:"amount" binding:"required" example:"1250.50"`
	Currency    string    `json:"currency" binding:"required" example:"RUB"`
	Category    string    `json:"category" binding:"required" example:"Продукты"`
	Description string    `json:"description" example:"Покупка в магазине"`
	OccurredAt  time.Time `json:"occurred_at" binding:"required" example:"2024-01-01T10:00:00Z"`
}

// Budget описывает бюджет.
type Budget struct {
	ID        string    `json:"id" example:"11111111-1111-1111-1111-111111111111"`
	Name      string    `json:"name" example:"Еда"`
	Amount    float64   `json:"amount" example:"10000"`
	Currency  string    `json:"currency" example:"RUB"`
	Period    string    `json:"period" example:"monthly"`
	StartDate time.Time `json:"start_date" example:"2024-01-01T00:00:00Z"`
	EndDate   time.Time `json:"end_date" example:"2024-01-31T23:59:59Z"`
	CreatedAt time.Time `json:"created_at" example:"2024-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2024-01-01T00:00:00Z"`
}

// CreateBudgetRequest описывает запрос на создание бюджета.
type CreateBudgetRequest struct {
	Name      string    `json:"name" binding:"required" example:"Еда"`
	Amount    float64   `json:"amount" binding:"required" example:"10000"`
	Currency  string    `json:"currency" binding:"required" example:"RUB"`
	Period    string    `json:"period" binding:"required" example:"monthly"`
	StartDate time.Time `json:"start_date" binding:"required" example:"2024-01-01T00:00:00Z"`
	EndDate   time.Time `json:"end_date" binding:"required" example:"2024-01-31T23:59:59Z"`
}

// UpdateBudgetRequest описывает запрос на обновление бюджета.
type UpdateBudgetRequest struct {
	Name      string    `json:"name" binding:"required" example:"Еда"`
	Amount    float64   `json:"amount" binding:"required" example:"10000"`
	Currency  string    `json:"currency" binding:"required" example:"RUB"`
	Period    string    `json:"period" binding:"required" example:"monthly"`
	StartDate time.Time `json:"start_date" binding:"required" example:"2024-01-01T00:00:00Z"`
	EndDate   time.Time `json:"end_date" binding:"required" example:"2024-01-31T23:59:59Z"`
}

// Report описывает отчет.
type Report struct {
	ID           string    `json:"id" example:"11111111-1111-1111-1111-111111111111"`
	Name         string    `json:"name" example:"Январь 2024"`
	Period       string    `json:"period" example:"2024-01"`
	GeneratedAt  time.Time `json:"generated_at" example:"2024-01-31T23:59:59Z"`
	TotalIncome  float64   `json:"total_income" example:"50000"`
	TotalExpense float64   `json:"total_expense" example:"30000"`
	Currency     string    `json:"currency" example:"RUB"`
}

// CreateReportRequest описывает запрос на создание отчета.
type CreateReportRequest struct {
	Name        string    `json:"name" binding:"required" example:"Январь 2024"`
	Period      string    `json:"period" binding:"required" example:"2024-01"`
	GeneratedAt time.Time `json:"generated_at" binding:"required" example:"2024-01-31T23:59:59Z"`
	Currency    string    `json:"currency" binding:"required" example:"RUB"`
}

// UpdateReportRequest описывает запрос на обновление отчета.
type UpdateReportRequest struct {
	Name        string    `json:"name" binding:"required" example:"Январь 2024"`
	Period      string    `json:"period" binding:"required" example:"2024-01"`
	GeneratedAt time.Time `json:"generated_at" binding:"required" example:"2024-01-31T23:59:59Z"`
	Currency    string    `json:"currency" binding:"required" example:"RUB"`
}

// ImportTransactionsRequest описывает импорт транзакций из CSV.
type ImportTransactionsRequest struct {
	CSVContent string `json:"csv_content" binding:"required" example:"account_id,amount,currency,category,description,occurred_at"`
	HasHeader  bool   `json:"has_header"`
}

// ImportTransactionsResponse описывает результат импорта.
type ImportTransactionsResponse struct {
	Imported int32 `json:"imported" example:"3"`
}

// ExportTransactionsResponse описывает экспорт в CSV.
type ExportTransactionsResponse struct {
	CSVContent string `json:"csv_content" example:"account_id,amount,currency,category,description,occurred_at"`
}

// DeleteResponse описывает ответ на удаление.
type DeleteResponse struct {
	Deleted bool `json:"deleted" example:"true"`
}
