package service

import (
	"context"
	"fmt"
	"time"

	"github.com/Deevins/final-task-course-2-go-lang/ledger/internal/model"
)

type ValidationService struct {
	next LedgerService
}

func NewValidationService(next LedgerService) *ValidationService {
	return &ValidationService{next: next}
}

func (s *ValidationService) CreateTransaction(ctx context.Context, tx model.Transaction) (model.Transaction, error) {
	if err := validateTransaction(tx, false); err != nil {
		return model.Transaction{}, err
	}
	return s.next.CreateTransaction(ctx, tx)
}

func (s *ValidationService) GetTransaction(ctx context.Context, id string) (model.Transaction, error) {
	if id == "" {
		return model.Transaction{}, fmt.Errorf("%w: transaction id is required", ErrValidation)
	}
	return s.next.GetTransaction(ctx, id)
}

func (s *ValidationService) UpdateTransaction(ctx context.Context, tx model.Transaction) (model.Transaction, error) {
	if err := validateTransaction(tx, true); err != nil {
		return model.Transaction{}, err
	}
	return s.next.UpdateTransaction(ctx, tx)
}

func (s *ValidationService) DeleteTransaction(ctx context.Context, id string) error {
	if id == "" {
		return fmt.Errorf("%w: transaction id is required", ErrValidation)
	}
	return s.next.DeleteTransaction(ctx, id)
}

func (s *ValidationService) ListTransactions(ctx context.Context, accountID string) []model.Transaction {
	return s.next.ListTransactions(ctx, accountID)
}

func (s *ValidationService) CreateBudget(ctx context.Context, budget model.Budget) (model.Budget, error) {
	if err := validateBudget(budget, false); err != nil {
		return model.Budget{}, err
	}
	return s.next.CreateBudget(ctx, budget)
}

func (s *ValidationService) GetBudget(ctx context.Context, accountID, id string) (model.Budget, error) {
	if accountID == "" {
		return model.Budget{}, fmt.Errorf("%w: account id is required", ErrValidation)
	}
	if id == "" {
		return model.Budget{}, fmt.Errorf("%w: budget id is required", ErrValidation)
	}
	return s.next.GetBudget(ctx, accountID, id)
}

func (s *ValidationService) UpdateBudget(ctx context.Context, budget model.Budget) (model.Budget, error) {
	if err := validateBudget(budget, true); err != nil {
		return model.Budget{}, err
	}
	return s.next.UpdateBudget(ctx, budget)
}

func (s *ValidationService) DeleteBudget(ctx context.Context, accountID, id string) error {
	if accountID == "" {
		return fmt.Errorf("%w: account id is required", ErrValidation)
	}
	if id == "" {
		return fmt.Errorf("%w: budget id is required", ErrValidation)
	}
	return s.next.DeleteBudget(ctx, accountID, id)
}

func (s *ValidationService) ListBudgets(ctx context.Context, accountID string) []model.Budget {
	return s.next.ListBudgets(ctx, accountID)
}

func (s *ValidationService) CreateReport(ctx context.Context, report model.Report) (model.Report, error) {
	if report.AccountID == "" {
		return model.Report{}, fmt.Errorf("%w: account id is required", ErrValidation)
	}
	if report.Period == "" {
		return model.Report{}, fmt.Errorf("%w: period is required", ErrValidation)
	}
	if report.Name == "" {
		return model.Report{}, fmt.Errorf("%w: report name is required", ErrValidation)
	}
	return s.next.CreateReport(ctx, report)
}

func (s *ValidationService) GetReport(ctx context.Context, accountID, id string) (model.Report, error) {
	if accountID == "" {
		return model.Report{}, fmt.Errorf("%w: account id is required", ErrValidation)
	}
	if id == "" {
		return model.Report{}, fmt.Errorf("%w: report id is required", ErrValidation)
	}
	return s.next.GetReport(ctx, accountID, id)
}

func (s *ValidationService) UpdateReport(ctx context.Context, report model.Report) (model.Report, error) {
	if report.ID == "" {
		return model.Report{}, fmt.Errorf("%w: report id is required", ErrValidation)
	}
	if report.AccountID == "" {
		return model.Report{}, fmt.Errorf("%w: account id is required", ErrValidation)
	}
	if report.Period == "" {
		return model.Report{}, fmt.Errorf("%w: period is required", ErrValidation)
	}
	if report.Name == "" {
		return model.Report{}, fmt.Errorf("%w: report name is required", ErrValidation)
	}
	return s.next.UpdateReport(ctx, report)
}

func (s *ValidationService) DeleteReport(ctx context.Context, accountID, id string) error {
	if accountID == "" {
		return fmt.Errorf("%w: account id is required", ErrValidation)
	}
	if id == "" {
		return fmt.Errorf("%w: report id is required", ErrValidation)
	}
	return s.next.DeleteReport(ctx, accountID, id)
}

func (s *ValidationService) ListReports(ctx context.Context, accountID string) []model.Report {
	return s.next.ListReports(ctx, accountID)
}

func (s *ValidationService) ImportTransactionsCSV(ctx context.Context, accountID string, csvContent []byte, hasHeader bool) (int, error) {
	if accountID == "" {
		return 0, fmt.Errorf("%w: account id is required", ErrValidation)
	}
	if len(csvContent) == 0 {
		return 0, fmt.Errorf("%w: csv content is required", ErrValidation)
	}
	return s.next.ImportTransactionsCSV(ctx, accountID, csvContent, hasHeader)
}

func (s *ValidationService) ExportTransactionsCSV(ctx context.Context, accountID string) ([]byte, error) {
	return s.next.ExportTransactionsCSV(ctx, accountID)
}

func validateTransaction(tx model.Transaction, requireID bool) error {
	if requireID && tx.ID == "" {
		return fmt.Errorf("%w: transaction id is required", ErrValidation)
	}
	if tx.AccountID == "" {
		return fmt.Errorf("%w: account id is required", ErrValidation)
	}
	if tx.Amount == 0 {
		return fmt.Errorf("%w: amount must be non-zero", ErrValidation)
	}
	if tx.Currency == "" {
		return fmt.Errorf("%w: currency is required", ErrValidation)
	}
	if tx.Category == "" {
		return fmt.Errorf("%w: category is required", ErrValidation)
	}
	return nil
}

func validateBudget(budget model.Budget, requireID bool) error {
	if requireID && budget.ID == "" {
		return fmt.Errorf("%w: budget id is required", ErrValidation)
	}
	if budget.AccountID == "" {
		return fmt.Errorf("%w: account id is required", ErrValidation)
	}
	if budget.Name == "" {
		return fmt.Errorf("%w: budget name is required", ErrValidation)
	}
	if budget.Amount <= 0 {
		return fmt.Errorf("%w: budget amount must be positive", ErrValidation)
	}
	if budget.Currency == "" {
		return fmt.Errorf("%w: currency is required", ErrValidation)
	}
	if budget.Period != "monthly" {
		return fmt.Errorf("%w: period must be monthly", ErrValidation)
	}
	if err := validateMonth(budget.Month); err != nil {
		return fmt.Errorf("%w: %v", ErrValidation, err)
	}
	return nil
}

func validateMonth(month time.Time) error {
	if month.IsZero() {
		return fmt.Errorf("month is required")
	}
	if month.Location() != time.UTC {
		month = month.UTC()
	}
	if month.Day() != 1 || month.Hour() != 0 || month.Minute() != 0 || month.Second() != 0 || month.Nanosecond() != 0 {
		return fmt.Errorf("month must be the first day of the month at 00:00:00Z")
	}
	return nil
}
