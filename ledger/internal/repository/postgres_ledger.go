package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/Deevins/final-task-course-2-go-lang/ledger/internal/model"
	"github.com/Deevins/final-task-course-2-go-lang/ledger/internal/storage"
)

type PostgresLedgerRepository struct {
	transactions *PostgresTransactionRepository
	budgets      *PostgresBudgetRepository
	reports      *storage.InMemoryLedgerStorage
}

func NewPostgresLedgerRepository(db *pgxpool.Pool) *PostgresLedgerRepository {
	return &PostgresLedgerRepository{
		transactions: NewPostgresTransactionRepository(db),
		budgets:      NewPostgresBudgetRepository(db),
		reports:      storage.NewInMemoryLedgerStorage(),
	}
}

func (r *PostgresLedgerRepository) CreateTransaction(tx model.Transaction) (model.Transaction, error) {
	return r.transactions.CreateTransaction(tx)
}

func (r *PostgresLedgerRepository) GetTransaction(id string) (model.Transaction, error) {
	return r.transactions.GetTransaction(id)
}

func (r *PostgresLedgerRepository) UpdateTransaction(tx model.Transaction) (model.Transaction, error) {
	return r.transactions.UpdateTransaction(tx)
}

func (r *PostgresLedgerRepository) DeleteTransaction(id string) error {
	return r.transactions.DeleteTransaction(id)
}

func (r *PostgresLedgerRepository) ListTransactions() []model.Transaction {
	return r.transactions.ListTransactions()
}

func (r *PostgresLedgerRepository) CreateBudget(budget model.Budget) (model.Budget, error) {
	return r.budgets.CreateBudget(budget)
}

func (r *PostgresLedgerRepository) GetBudget(id string) (model.Budget, error) {
	return r.budgets.GetBudget(id)
}

func (r *PostgresLedgerRepository) UpdateBudget(budget model.Budget) (model.Budget, error) {
	return r.budgets.UpdateBudget(budget)
}

func (r *PostgresLedgerRepository) DeleteBudget(id string) error {
	return r.budgets.DeleteBudget(id)
}

func (r *PostgresLedgerRepository) ListBudgets() []model.Budget {
	return r.budgets.ListBudgets()
}

func (r *PostgresLedgerRepository) CreateReport(report model.Report) (model.Report, error) {
	return r.reports.CreateReport(report), nil
}

func (r *PostgresLedgerRepository) GetReport(id string) (model.Report, error) {
	return r.reports.GetReport(id)
}

func (r *PostgresLedgerRepository) UpdateReport(report model.Report) (model.Report, error) {
	return r.reports.UpdateReport(report)
}

func (r *PostgresLedgerRepository) DeleteReport(id string) error {
	return r.reports.DeleteReport(id)
}

func (r *PostgresLedgerRepository) ListReports() []model.Report {
	return r.reports.ListReports()
}
