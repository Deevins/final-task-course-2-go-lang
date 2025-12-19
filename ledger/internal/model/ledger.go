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
	AccountID string
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
	AccountID    string
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
