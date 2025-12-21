package repository

import (
	"context"

	"github.com/Deevins/final-task-course-2-go-lang/ledger/internal/model"
	"github.com/Deevins/final-task-course-2-go-lang/ledger/internal/storage"
)

type LedgerRepository interface {
	CreateTransaction(ctx context.Context, tx model.Transaction) (model.Transaction, error)
	GetTransaction(ctx context.Context, id string) (model.Transaction, error)
	UpdateTransaction(ctx context.Context, tx model.Transaction) (model.Transaction, error)
	DeleteTransaction(ctx context.Context, id string) error
	ListTransactions(ctx context.Context) []model.Transaction

	CreateBudget(ctx context.Context, budget model.Budget) (model.Budget, error)
	GetBudget(ctx context.Context, accountID, id string) (model.Budget, error)
	UpdateBudget(ctx context.Context, budget model.Budget) (model.Budget, error)
	DeleteBudget(ctx context.Context, accountID, id string) error
	ListBudgets(ctx context.Context, accountID string) []model.Budget

	CreateReport(ctx context.Context, report model.Report) (model.Report, error)
	GetReport(ctx context.Context, accountID, id string) (model.Report, error)
	UpdateReport(ctx context.Context, report model.Report) (model.Report, error)
	DeleteReport(ctx context.Context, accountID, id string) error
	ListReports(ctx context.Context, accountID string) []model.Report
}

type InMemoryLedgerRepository struct {
	store *storage.InMemoryLedgerStorage
}

func NewInMemoryLedgerRepository(store *storage.InMemoryLedgerStorage) *InMemoryLedgerRepository {
	return &InMemoryLedgerRepository{store: store}
}

func (r *InMemoryLedgerRepository) CreateTransaction(ctx context.Context, tx model.Transaction) (model.Transaction, error) {
	return r.store.CreateTransaction(tx), nil
}

func (r *InMemoryLedgerRepository) GetTransaction(ctx context.Context, id string) (model.Transaction, error) {
	return r.store.GetTransaction(id)
}

func (r *InMemoryLedgerRepository) UpdateTransaction(ctx context.Context, tx model.Transaction) (model.Transaction, error) {
	return r.store.UpdateTransaction(tx)
}

func (r *InMemoryLedgerRepository) DeleteTransaction(ctx context.Context, id string) error {
	return r.store.DeleteTransaction(id)
}

func (r *InMemoryLedgerRepository) ListTransactions(ctx context.Context) []model.Transaction {
	return r.store.ListTransactions()
}

func (r *InMemoryLedgerRepository) CreateBudget(ctx context.Context, budget model.Budget) (model.Budget, error) {
	return r.store.CreateBudget(budget), nil
}

func (r *InMemoryLedgerRepository) GetBudget(ctx context.Context, accountID, id string) (model.Budget, error) {
	budget, err := r.store.GetBudget(id)
	if err != nil {
		return model.Budget{}, err
	}
	if budget.AccountID != accountID {
		return model.Budget{}, storage.ErrNotFound
	}
	return budget, nil
}

func (r *InMemoryLedgerRepository) UpdateBudget(ctx context.Context, budget model.Budget) (model.Budget, error) {
	return r.store.UpdateBudget(budget)
}

func (r *InMemoryLedgerRepository) DeleteBudget(ctx context.Context, accountID, id string) error {
	budget, err := r.store.GetBudget(id)
	if err != nil {
		return err
	}
	if budget.AccountID != accountID {
		return storage.ErrNotFound
	}
	return r.store.DeleteBudget(id)
}

func (r *InMemoryLedgerRepository) ListBudgets(ctx context.Context, accountID string) []model.Budget {
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

func (r *InMemoryLedgerRepository) CreateReport(ctx context.Context, report model.Report) (model.Report, error) {
	return r.store.CreateReport(report), nil
}

func (r *InMemoryLedgerRepository) GetReport(ctx context.Context, accountID, id string) (model.Report, error) {
	report, err := r.store.GetReport(id)
	if err != nil {
		return model.Report{}, err
	}
	if report.AccountID != accountID {
		return model.Report{}, storage.ErrNotFound
	}
	return report, nil
}

func (r *InMemoryLedgerRepository) UpdateReport(ctx context.Context, report model.Report) (model.Report, error) {
	return r.store.UpdateReport(report)
}

func (r *InMemoryLedgerRepository) DeleteReport(ctx context.Context, accountID, id string) error {
	report, err := r.store.GetReport(id)
	if err != nil {
		return err
	}
	if report.AccountID != accountID {
		return storage.ErrNotFound
	}
	return r.store.DeleteReport(id)
}

func (r *InMemoryLedgerRepository) ListReports(ctx context.Context, accountID string) []model.Report {
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
