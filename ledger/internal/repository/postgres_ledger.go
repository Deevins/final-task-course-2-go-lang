package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/Deevins/final-task-course-2-go-lang/ledger/internal/model"
)

type PostgresLedgerRepository struct {
	transactions *PostgresTransactionRepository
	budgets      *PostgresBudgetRepository
	reports      *PostgresReportRepository
}

func NewPostgresLedgerRepository(db *pgxpool.Pool) *PostgresLedgerRepository {
	return &PostgresLedgerRepository{
		transactions: NewPostgresTransactionRepository(db),
		budgets:      NewPostgresBudgetRepository(db),
		reports:      NewPostgresReportRepository(db),
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

func (r *PostgresLedgerRepository) GetBudget(accountID, id string) (model.Budget, error) {
	return r.budgets.GetBudget(accountID, id)
}

func (r *PostgresLedgerRepository) UpdateBudget(budget model.Budget) (model.Budget, error) {
	return r.budgets.UpdateBudget(budget)
}

func (r *PostgresLedgerRepository) DeleteBudget(accountID, id string) error {
	return r.budgets.DeleteBudget(accountID, id)
}

func (r *PostgresLedgerRepository) ListBudgets(accountID string) []model.Budget {
	return r.budgets.ListBudgets(accountID)
}

func (r *PostgresLedgerRepository) CreateReport(report model.Report) (model.Report, error) {
	return r.reports.CreateReport(report)
}

func (r *PostgresLedgerRepository) GetReport(accountID, id string) (model.Report, error) {
	return r.reports.GetReport(accountID, id)
}

func (r *PostgresLedgerRepository) UpdateReport(report model.Report) (model.Report, error) {
	return r.reports.UpdateReport(report)
}

func (r *PostgresLedgerRepository) DeleteReport(accountID, id string) error {
	return r.reports.DeleteReport(accountID, id)
}

func (r *PostgresLedgerRepository) ListReports(accountID string) []model.Report {
	return r.reports.ListReports(accountID)
}
