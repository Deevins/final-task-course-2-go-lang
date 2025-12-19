package model

import "time"

type Transaction struct {
	ID          string
	AccountID   string
	Amount      float64
	Currency    string
	Category    string
	Description string
	OccurredAt  time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Budget struct {
	ID        string
	Name      string
	Amount    float64
	Currency  string
	Period    string
	StartDate time.Time
	EndDate   time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Report struct {
	ID           string
	Name         string
	Period       string
	GeneratedAt  time.Time
	TotalIncome  float64
	TotalExpense float64
	Currency     string
	Categories   []ReportCategory
}

type ReportCategory struct {
	Category           string
	TotalExpense       float64
	BudgetAmount       float64
	BudgetUsagePercent float64
}

type TransactionCSVRow struct {
	AccountID   string
	Amount      float64
	Currency    string
	Category    string
	Description string
	OccurredAt  time.Time
}

type ReportSheetSummary struct {
	ReportID     string    `json:"report_id"`
	Name         string    `json:"name"`
	Period       string    `json:"period"`
	GeneratedAt  time.Time `json:"generated_at"`
	TotalIncome  float64   `json:"total_income"`
	TotalExpense float64   `json:"total_expense"`
	Currency     string    `json:"currency"`
}

type ReportSheetCategory struct {
	Category           string  `json:"category"`
	TotalExpense       float64 `json:"total_expense"`
	BudgetAmount       float64 `json:"budget_amount"`
	BudgetUsagePercent float64 `json:"budget_usage_percent"`
}

type ReportSheet struct {
	Summary    ReportSheetSummary    `json:"summary"`
	Categories []ReportSheetCategory `json:"categories"`
}
