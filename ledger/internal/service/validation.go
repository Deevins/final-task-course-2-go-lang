package service

import (
	"fmt"

	"github.com/Deevins/final-task-course-2-go-lang/ledger/internal/model"
)

type ValidationService struct {
	next LedgerService
}

func NewValidationService(next LedgerService) *ValidationService {
	return &ValidationService{next: next}
}

func (s *ValidationService) CreateTransaction(tx model.Transaction) (model.Transaction, error) {
	if err := validateTransaction(tx, false); err != nil {
		return model.Transaction{}, err
	}
	return s.next.CreateTransaction(tx)
}

func (s *ValidationService) GetTransaction(id string) (model.Transaction, error) {
	if id == "" {
		return model.Transaction{}, fmt.Errorf("%w: transaction id is required", ErrValidation)
	}
	return s.next.GetTransaction(id)
}

func (s *ValidationService) UpdateTransaction(tx model.Transaction) (model.Transaction, error) {
	if err := validateTransaction(tx, true); err != nil {
		return model.Transaction{}, err
	}
	return s.next.UpdateTransaction(tx)
}

func (s *ValidationService) DeleteTransaction(id string) error {
	if id == "" {
		return fmt.Errorf("%w: transaction id is required", ErrValidation)
	}
	return s.next.DeleteTransaction(id)
}

func (s *ValidationService) ListTransactions(accountID string) []model.Transaction {
	return s.next.ListTransactions(accountID)
}

func (s *ValidationService) CreateBudget(budget model.Budget) (model.Budget, error) {
	if err := validateBudget(budget, false); err != nil {
		return model.Budget{}, err
	}
	return s.next.CreateBudget(budget)
}

func (s *ValidationService) GetBudget(accountID, id string) (model.Budget, error) {
	if accountID == "" {
		return model.Budget{}, fmt.Errorf("%w: account id is required", ErrValidation)
	}
	if id == "" {
		return model.Budget{}, fmt.Errorf("%w: budget id is required", ErrValidation)
	}
	return s.next.GetBudget(accountID, id)
}

func (s *ValidationService) UpdateBudget(budget model.Budget) (model.Budget, error) {
	if err := validateBudget(budget, true); err != nil {
		return model.Budget{}, err
	}
	return s.next.UpdateBudget(budget)
}

func (s *ValidationService) DeleteBudget(accountID, id string) error {
	if accountID == "" {
		return fmt.Errorf("%w: account id is required", ErrValidation)
	}
	if id == "" {
		return fmt.Errorf("%w: budget id is required", ErrValidation)
	}
	return s.next.DeleteBudget(accountID, id)
}

func (s *ValidationService) ListBudgets(accountID string) []model.Budget {
	return s.next.ListBudgets(accountID)
}

func (s *ValidationService) CreateReport(report model.Report) (model.Report, error) {
	if report.AccountID == "" {
		return model.Report{}, fmt.Errorf("%w: account id is required", ErrValidation)
	}
	if report.Period == "" {
		return model.Report{}, fmt.Errorf("%w: period is required", ErrValidation)
	}
	if report.Name == "" {
		return model.Report{}, fmt.Errorf("%w: report name is required", ErrValidation)
	}
	return s.next.CreateReport(report)
}

func (s *ValidationService) GetReport(accountID, id string) (model.Report, error) {
	if accountID == "" {
		return model.Report{}, fmt.Errorf("%w: account id is required", ErrValidation)
	}
	if id == "" {
		return model.Report{}, fmt.Errorf("%w: report id is required", ErrValidation)
	}
	return s.next.GetReport(accountID, id)
}

func (s *ValidationService) UpdateReport(report model.Report) (model.Report, error) {
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
	return s.next.UpdateReport(report)
}

func (s *ValidationService) DeleteReport(accountID, id string) error {
	if accountID == "" {
		return fmt.Errorf("%w: account id is required", ErrValidation)
	}
	if id == "" {
		return fmt.Errorf("%w: report id is required", ErrValidation)
	}
	return s.next.DeleteReport(accountID, id)
}

func (s *ValidationService) ListReports(accountID string) []model.Report {
	return s.next.ListReports(accountID)
}

func (s *ValidationService) ImportTransactionsCSV(csvContent []byte, hasHeader bool) (int, error) {
	if len(csvContent) == 0 {
		return 0, fmt.Errorf("%w: csv content is required", ErrValidation)
	}
	return s.next.ImportTransactionsCSV(csvContent, hasHeader)
}

func (s *ValidationService) ExportTransactionsCSV(accountID string) ([]byte, error) {
	return s.next.ExportTransactionsCSV(accountID)
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
	if !budget.StartDate.IsZero() && !budget.EndDate.IsZero() {
		if budget.EndDate.Before(budget.StartDate) {
			return fmt.Errorf("%w: end date must be after start date", ErrValidation)
		}
	}
	return nil
}
