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
	GetBudget(id string) (model.Budget, error)
	UpdateBudget(budget model.Budget) (model.Budget, error)
	DeleteBudget(id string) error
	ListBudgets() []model.Budget

	CreateReport(report model.Report) (model.Report, error)
	GetReport(id string) (model.Report, error)
	UpdateReport(report model.Report) (model.Report, error)
	DeleteReport(id string) error
	ListReports() []model.Report
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

func (r *InMemoryLedgerRepository) GetBudget(id string) (model.Budget, error) {
	return r.store.GetBudget(id)
}

func (r *InMemoryLedgerRepository) UpdateBudget(budget model.Budget) (model.Budget, error) {
	return r.store.UpdateBudget(budget)
}

func (r *InMemoryLedgerRepository) DeleteBudget(id string) error {
	return r.store.DeleteBudget(id)
}

func (r *InMemoryLedgerRepository) ListBudgets() []model.Budget {
	return r.store.ListBudgets()
}

func (r *InMemoryLedgerRepository) CreateReport(report model.Report) (model.Report, error) {
	return r.store.CreateReport(report), nil
}

func (r *InMemoryLedgerRepository) GetReport(id string) (model.Report, error) {
	return r.store.GetReport(id)
}

func (r *InMemoryLedgerRepository) UpdateReport(report model.Report) (model.Report, error) {
	return r.store.UpdateReport(report)
}

func (r *InMemoryLedgerRepository) DeleteReport(id string) error {
	return r.store.DeleteReport(id)
}

func (r *InMemoryLedgerRepository) ListReports() []model.Report {
	return r.store.ListReports()
}
