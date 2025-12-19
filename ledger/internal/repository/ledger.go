package repository

import (
	"github.com/Deevins/final-task-course-2-go-lang/ledger/internal/model"
	"github.com/Deevins/final-task-course-2-go-lang/ledger/internal/storage"
)

type LedgerRepository interface {
	CreateTransaction(tx model.Transaction) (model.Transaction, error)
	GetTransaction(id string) (model.Transaction, error)
	UpdateTransaction(tx model.Transaction) (model.Transaction, error)
	DeleteTransaction(id string) error
	ListTransactions() []model.Transaction

	CreateBudget(budget model.Budget) (model.Budget, error)
	GetBudget(accountID, id string) (model.Budget, error)
	UpdateBudget(budget model.Budget) (model.Budget, error)
	DeleteBudget(accountID, id string) error
	ListBudgets(accountID string) []model.Budget

	CreateReport(report model.Report) (model.Report, error)
	GetReport(accountID, id string) (model.Report, error)
	UpdateReport(report model.Report) (model.Report, error)
	DeleteReport(accountID, id string) error
	ListReports(accountID string) []model.Report
}

type InMemoryLedgerRepository struct {
	store *storage.InMemoryLedgerStorage
}

func NewInMemoryLedgerRepository(store *storage.InMemoryLedgerStorage) *InMemoryLedgerRepository {
	return &InMemoryLedgerRepository{store: store}
}

func (r *InMemoryLedgerRepository) CreateTransaction(tx model.Transaction) (model.Transaction, error) {
	return r.store.CreateTransaction(tx), nil
}

func (r *InMemoryLedgerRepository) GetTransaction(id string) (model.Transaction, error) {
	return r.store.GetTransaction(id)
}

func (r *InMemoryLedgerRepository) UpdateTransaction(tx model.Transaction) (model.Transaction, error) {
	return r.store.UpdateTransaction(tx)
}

func (r *InMemoryLedgerRepository) DeleteTransaction(id string) error {
	return r.store.DeleteTransaction(id)
}

func (r *InMemoryLedgerRepository) ListTransactions() []model.Transaction {
	return r.store.ListTransactions()
}

func (r *InMemoryLedgerRepository) CreateBudget(budget model.Budget) (model.Budget, error) {
	return r.store.CreateBudget(budget), nil
}

func (r *InMemoryLedgerRepository) GetBudget(accountID, id string) (model.Budget, error) {
	budget, err := r.store.GetBudget(id)
	if err != nil {
		return model.Budget{}, err
	}
	if budget.AccountID != accountID {
		return model.Budget{}, storage.ErrNotFound
	}
	return budget, nil
}

func (r *InMemoryLedgerRepository) UpdateBudget(budget model.Budget) (model.Budget, error) {
	return r.store.UpdateBudget(budget)
}

func (r *InMemoryLedgerRepository) DeleteBudget(accountID, id string) error {
	budget, err := r.store.GetBudget(id)
	if err != nil {
		return err
	}
	if budget.AccountID != accountID {
		return storage.ErrNotFound
	}
	return r.store.DeleteBudget(id)
}

func (r *InMemoryLedgerRepository) ListBudgets(accountID string) []model.Budget {
	items := r.store.ListBudgets()
	if accountID == "" {
		return nil
	}
	filtered := make([]model.Budget, 0, len(items))
	for _, budget := range items {
		if budget.AccountID == accountID {
			filtered = append(filtered, budget)
		}
	}
	return filtered
}

func (r *InMemoryLedgerRepository) CreateReport(report model.Report) (model.Report, error) {
	return r.store.CreateReport(report), nil
}

func (r *InMemoryLedgerRepository) GetReport(accountID, id string) (model.Report, error) {
	report, err := r.store.GetReport(id)
	if err != nil {
		return model.Report{}, err
	}
	if report.AccountID != accountID {
		return model.Report{}, storage.ErrNotFound
	}
	return report, nil
}

func (r *InMemoryLedgerRepository) UpdateReport(report model.Report) (model.Report, error) {
	return r.store.UpdateReport(report)
}

func (r *InMemoryLedgerRepository) DeleteReport(accountID, id string) error {
	report, err := r.store.GetReport(id)
	if err != nil {
		return err
	}
	if report.AccountID != accountID {
		return storage.ErrNotFound
	}
	return r.store.DeleteReport(id)
}

func (r *InMemoryLedgerRepository) ListReports(accountID string) []model.Report {
	items := r.store.ListReports()
	if accountID == "" {
		return nil
	}
	filtered := make([]model.Report, 0, len(items))
	for _, report := range items {
		if report.AccountID == accountID {
			filtered = append(filtered, report)
		}
	}
	return filtered
}
