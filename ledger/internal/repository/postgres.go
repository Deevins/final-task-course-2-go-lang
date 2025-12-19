package repository

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/Deevins/final-task-course-2-go-lang/ledger/internal/model"
	"github.com/Deevins/final-task-course-2-go-lang/ledger/internal/storage"
)

type TransactionRepository interface {
	CreateTransaction(ctx context.Context, tx model.Transaction) (model.Transaction, error)
	GetTransaction(ctx context.Context, id string) (model.Transaction, error)
	UpdateTransaction(ctx context.Context, tx model.Transaction) (model.Transaction, error)
	DeleteTransaction(ctx context.Context, id string) error
	ListTransactions(ctx context.Context) []model.Transaction
}

type BudgetRepository interface {
	CreateBudget(ctx context.Context, budget model.Budget) (model.Budget, error)
	GetBudget(ctx context.Context, accountID, id string) (model.Budget, error)
	UpdateBudget(ctx context.Context, budget model.Budget) (model.Budget, error)
	DeleteBudget(ctx context.Context, accountID, id string) error
	ListBudgets(ctx context.Context, accountID string) []model.Budget
}

type ReportRepository interface {
	CreateReport(ctx context.Context, report model.Report) (model.Report, error)
	GetReport(ctx context.Context, accountID, id string) (model.Report, error)
	UpdateReport(ctx context.Context, report model.Report) (model.Report, error)
	DeleteReport(ctx context.Context, accountID, id string) error
	ListReports(ctx context.Context, accountID string) []model.Report
}

type PostgresTransactionRepository struct {
	db *pgxpool.Pool
}

func NewPostgresTransactionRepository(db *pgxpool.Pool) *PostgresTransactionRepository {
	return &PostgresTransactionRepository{db: db}
}

func (r *PostgresTransactionRepository) CreateTransaction(ctx context.Context, tx model.Transaction) (model.Transaction, error) {
	const query = `
		INSERT INTO transactions (id, account_id, amount, currency, category, description, occurred_at, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err := r.db.Exec(ctx, query, tx.ID, tx.AccountID, tx.Amount, tx.Currency, tx.Category, tx.Description, tx.OccurredAt, tx.CreatedAt, tx.UpdatedAt)
	if err != nil {
		return model.Transaction{}, err
	}
	return tx, nil
}

func (r *PostgresTransactionRepository) GetTransaction(ctx context.Context, id string) (model.Transaction, error) {
	const query = `
		SELECT id, account_id, amount, currency, category, description, occurred_at, created_at, updated_at
		FROM transactions
		WHERE id = $1`
	var tx model.Transaction
	err := r.db.QueryRow(ctx, query, id).Scan(
		&tx.ID,
		&tx.AccountID,
		&tx.Amount,
		&tx.Currency,
		&tx.Category,
		&tx.Description,
		&tx.OccurredAt,
		&tx.CreatedAt,
		&tx.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Transaction{}, storage.ErrNotFound
		}
		return model.Transaction{}, err
	}
	return tx, nil
}

func (r *PostgresTransactionRepository) UpdateTransaction(ctx context.Context, tx model.Transaction) (model.Transaction, error) {
	const query = `
		UPDATE transactions
		SET account_id = $2, amount = $3, currency = $4, category = $5, description = $6, occurred_at = $7, created_at = $8, updated_at = $9
		WHERE id = $1`
	result, err := r.db.Exec(ctx, query, tx.ID, tx.AccountID, tx.Amount, tx.Currency, tx.Category, tx.Description, tx.OccurredAt, tx.CreatedAt, tx.UpdatedAt)
	if err != nil {
		return model.Transaction{}, err
	}
	if result.RowsAffected() == 0 {
		return model.Transaction{}, storage.ErrNotFound
	}
	return tx, nil
}

func (r *PostgresTransactionRepository) DeleteTransaction(ctx context.Context, id string) error {
	const query = `DELETE FROM transactions WHERE id = $1`
	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return storage.ErrNotFound
	}
	return nil
}

func (r *PostgresTransactionRepository) ListTransactions(ctx context.Context) []model.Transaction {
	const query = `
		SELECT id, account_id, amount, currency, category, description, occurred_at, created_at, updated_at
		FROM transactions`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil
	}
	defer rows.Close()

	items := []model.Transaction{}
	for rows.Next() {
		var tx model.Transaction
		if err := rows.Scan(
			&tx.ID,
			&tx.AccountID,
			&tx.Amount,
			&tx.Currency,
			&tx.Category,
			&tx.Description,
			&tx.OccurredAt,
			&tx.CreatedAt,
			&tx.UpdatedAt,
		); err != nil {
			return nil
		}
		items = append(items, tx)
	}
	if err := rows.Err(); err != nil {
		return nil
	}
	return items
}

type PostgresBudgetRepository struct {
	db *pgxpool.Pool
}

func NewPostgresBudgetRepository(db *pgxpool.Pool) *PostgresBudgetRepository {
	return &PostgresBudgetRepository{db: db}
}

func (r *PostgresBudgetRepository) CreateBudget(ctx context.Context, budget model.Budget) (model.Budget, error) {
	const query = `
		INSERT INTO budgets (id, account_id, name, amount, currency, period, start_date, end_date, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	_, err := r.db.Exec(ctx, query, budget.ID, budget.AccountID, budget.Name, budget.Amount, budget.Currency, budget.Period, budget.StartDate, budget.EndDate, budget.CreatedAt, budget.UpdatedAt)
	if err != nil {
		return model.Budget{}, err
	}
	return budget, nil
}

func (r *PostgresBudgetRepository) GetBudget(ctx context.Context, accountID, id string) (model.Budget, error) {
	const query = `
		SELECT id, account_id, name, amount, currency, period, start_date, end_date, created_at, updated_at
		FROM budgets
		WHERE id = $1 AND account_id = $2`
	var budget model.Budget
	err := r.db.QueryRow(ctx, query, id, accountID).Scan(
		&budget.ID,
		&budget.AccountID,
		&budget.Name,
		&budget.Amount,
		&budget.Currency,
		&budget.Period,
		&budget.StartDate,
		&budget.EndDate,
		&budget.CreatedAt,
		&budget.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Budget{}, storage.ErrNotFound
		}
		return model.Budget{}, err
	}
	return budget, nil
}

func (r *PostgresBudgetRepository) UpdateBudget(ctx context.Context, budget model.Budget) (model.Budget, error) {
	const query = `
		UPDATE budgets
		SET name = $3, amount = $4, currency = $5, period = $6, start_date = $7, end_date = $8, created_at = $9, updated_at = $10
		WHERE id = $1 AND account_id = $2`
	result, err := r.db.Exec(ctx, query, budget.ID, budget.AccountID, budget.Name, budget.Amount, budget.Currency, budget.Period, budget.StartDate, budget.EndDate, budget.CreatedAt, budget.UpdatedAt)
	if err != nil {
		return model.Budget{}, err
	}
	if result.RowsAffected() == 0 {
		return model.Budget{}, storage.ErrNotFound
	}
	return budget, nil
}

func (r *PostgresBudgetRepository) DeleteBudget(ctx context.Context, accountID, id string) error {
	const query = `DELETE FROM budgets WHERE id = $1 AND account_id = $2`
	result, err := r.db.Exec(ctx, query, id, accountID)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return storage.ErrNotFound
	}
	return nil
}

func (r *PostgresBudgetRepository) ListBudgets(ctx context.Context, accountID string) []model.Budget {
	const query = `
		SELECT id, account_id, name, amount, currency, period, start_date, end_date, created_at, updated_at
		FROM budgets
		WHERE account_id = $1`
	rows, err := r.db.Query(ctx, query, accountID)
	if err != nil {
		return nil
	}
	defer rows.Close()

	items := []model.Budget{}
	for rows.Next() {
		var budget model.Budget
		if err := rows.Scan(
			&budget.ID,
			&budget.AccountID,
			&budget.Name,
			&budget.Amount,
			&budget.Currency,
			&budget.Period,
			&budget.StartDate,
			&budget.EndDate,
			&budget.CreatedAt,
			&budget.UpdatedAt,
		); err != nil {
			return nil
		}
		items = append(items, budget)
	}
	if err := rows.Err(); err != nil {
		return nil
	}
	return items
}

type PostgresReportRepository struct {
	db *pgxpool.Pool
}

func NewPostgresReportRepository(db *pgxpool.Pool) *PostgresReportRepository {
	return &PostgresReportRepository{db: db}
}

func (r *PostgresReportRepository) CreateReport(ctx context.Context, report model.Report) (model.Report, error) {
	const query = `
		INSERT INTO reports (id, account_id, name, period, generated_at, total_income, total_expense, currency, categories)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	categories, err := json.Marshal(report.Categories)
	if err != nil {
		return model.Report{}, err
	}
	_, err = r.db.Exec(
		ctx,
		query,
		report.ID,
		report.AccountID,
		report.Name,
		report.Period,
		report.GeneratedAt,
		report.TotalIncome,
		report.TotalExpense,
		report.Currency,
		categories,
	)
	if err != nil {
		return model.Report{}, err
	}
	return report, nil
}

func (r *PostgresReportRepository) GetReport(ctx context.Context, accountID, id string) (model.Report, error) {
	const query = `
		SELECT id, account_id, name, period, generated_at, total_income, total_expense, currency, categories
		FROM reports
		WHERE id = $1 AND account_id = $2`
	var report model.Report
	var categories []byte
	err := r.db.QueryRow(ctx, query, id, accountID).Scan(
		&report.ID,
		&report.AccountID,
		&report.Name,
		&report.Period,
		&report.GeneratedAt,
		&report.TotalIncome,
		&report.TotalExpense,
		&report.Currency,
		&categories,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Report{}, storage.ErrNotFound
		}
		return model.Report{}, err
	}
	if len(categories) > 0 {
		if err := json.Unmarshal(categories, &report.Categories); err != nil {
			return model.Report{}, err
		}
	}
	return report, nil
}

func (r *PostgresReportRepository) UpdateReport(ctx context.Context, report model.Report) (model.Report, error) {
	const query = `
		UPDATE reports
		SET name = $3, period = $4, generated_at = $5, total_income = $6, total_expense = $7, currency = $8, categories = $9
		WHERE id = $1 AND account_id = $2`
	categories, err := json.Marshal(report.Categories)
	if err != nil {
		return model.Report{}, err
	}
	result, err := r.db.Exec(
		ctx,
		query,
		report.ID,
		report.AccountID,
		report.Name,
		report.Period,
		report.GeneratedAt,
		report.TotalIncome,
		report.TotalExpense,
		report.Currency,
		categories,
	)
	if err != nil {
		return model.Report{}, err
	}
	if result.RowsAffected() == 0 {
		return model.Report{}, storage.ErrNotFound
	}
	return report, nil
}

func (r *PostgresReportRepository) DeleteReport(ctx context.Context, accountID, id string) error {
	const query = `DELETE FROM reports WHERE id = $1 AND account_id = $2`
	result, err := r.db.Exec(ctx, query, id, accountID)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return storage.ErrNotFound
	}
	return nil
}

func (r *PostgresReportRepository) ListReports(ctx context.Context, accountID string) []model.Report {
	const query = `
		SELECT id, account_id, name, period, generated_at, total_income, total_expense, currency, categories
		FROM reports
		WHERE account_id = $1`
	rows, err := r.db.Query(ctx, query, accountID)
	if err != nil {
		return nil
	}
	defer rows.Close()

	items := []model.Report{}
	for rows.Next() {
		var report model.Report
		var categories []byte
		if err := rows.Scan(
			&report.ID,
			&report.AccountID,
			&report.Name,
			&report.Period,
			&report.GeneratedAt,
			&report.TotalIncome,
			&report.TotalExpense,
			&report.Currency,
			&categories,
		); err != nil {
			return nil
		}
		if len(categories) > 0 {
			if err := json.Unmarshal(categories, &report.Categories); err != nil {
				return nil
			}
		}
		items = append(items, report)
	}
	if err := rows.Err(); err != nil {
		return nil
	}
	return items
}
