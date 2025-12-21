package repository

import (
	"context"

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

func (r *PostgresLedgerRepository) CreateTransaction(ctx context.Context, tx model.Transaction) (model.Transaction, error) {
	return r.transactions.CreateTransaction(ctx, tx)
}

func (r *PostgresLedgerRepository) GetTransaction(ctx context.Context, id string) (model.Transaction, error) {
	return r.transactions.GetTransaction(ctx, id)
}

func (r *PostgresLedgerRepository) UpdateTransaction(ctx context.Context, tx model.Transaction) (model.Transaction, error) {
	return r.transactions.UpdateTransaction(ctx, tx)
}

func (r *PostgresLedgerRepository) DeleteTransaction(ctx context.Context, id string) error {
	return r.transactions.DeleteTransaction(ctx, id)
}

func (r *PostgresLedgerRepository) ListTransactions(ctx context.Context) []model.Transaction {
	return r.transactions.ListTransactions(ctx)
}

func (r *PostgresLedgerRepository) CreateBudget(ctx context.Context, budget model.Budget) (model.Budget, error) {
	return r.budgets.CreateBudget(ctx, budget)
}

func (r *PostgresLedgerRepository) GetBudget(ctx context.Context, accountID, id string) (model.Budget, error) {
	return r.budgets.GetBudget(ctx, accountID, id)
}

func (r *PostgresLedgerRepository) UpdateBudget(ctx context.Context, budget model.Budget) (model.Budget, error) {
	return r.budgets.UpdateBudget(ctx, budget)
}

func (r *PostgresLedgerRepository) DeleteBudget(ctx context.Context, accountID, id string) error {
	return r.budgets.DeleteBudget(ctx, accountID, id)
}

func (r *PostgresLedgerRepository) ListBudgets(ctx context.Context, accountID string) []model.Budget {
	return r.budgets.ListBudgets(ctx, accountID)
}

func (r *PostgresLedgerRepository) CreateReport(ctx context.Context, report model.Report) (model.Report, error) {
	return r.reports.CreateReport(ctx, report)
}

func (r *PostgresLedgerRepository) GetReport(ctx context.Context, accountID, id string) (model.Report, error) {
	return r.reports.GetReport(ctx, accountID, id)
}

func (r *PostgresLedgerRepository) UpdateReport(ctx context.Context, report model.Report) (model.Report, error) {
	return r.reports.UpdateReport(ctx, report)
}

func (r *PostgresLedgerRepository) DeleteReport(ctx context.Context, accountID, id string) error {
	return r.reports.DeleteReport(ctx, accountID, id)
}

func (r *PostgresLedgerRepository) ListReports(ctx context.Context, accountID string) []model.Report {
	return r.reports.ListReports(ctx, accountID)
}
