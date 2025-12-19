package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/Deevins/final-task-course-2-go-lang/ledger/internal/model"
	"github.com/Deevins/final-task-course-2-go-lang/ledger/internal/storage"
)

type TransactionRepository interface {
	CreateTransaction(tx model.Transaction) (model.Transaction, error)
	GetTransaction(id string) (model.Transaction, error)
	UpdateTransaction(tx model.Transaction) (model.Transaction, error)
	DeleteTransaction(id string) error
	ListTransactions() []model.Transaction
}

type BudgetRepository interface {
	CreateBudget(budget model.Budget) (model.Budget, error)
	GetBudget(id string) (model.Budget, error)
	UpdateBudget(budget model.Budget) (model.Budget, error)
	DeleteBudget(id string) error
	ListBudgets() []model.Budget
}

type PostgresTransactionRepository struct {
	db *pgxpool.Pool
}

func NewPostgresTransactionRepository(db *pgxpool.Pool) *PostgresTransactionRepository {
	return &PostgresTransactionRepository{db: db}
}

func (r *PostgresTransactionRepository) CreateTransaction(tx model.Transaction) (model.Transaction, error) {
	const query = `
		INSERT INTO transactions (id, account_id, amount, currency, category, description, occurred_at, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err := r.db.Exec(context.Background(), query, tx.ID, tx.AccountID, tx.Amount, tx.Currency, tx.Category, tx.Description, tx.OccurredAt, tx.CreatedAt, tx.UpdatedAt)
	if err != nil {
		return model.Transaction{}, err
	}
	return tx, nil
}

func (r *PostgresTransactionRepository) GetTransaction(id string) (model.Transaction, error) {
	const query = `
		SELECT id, account_id, amount, currency, category, description, occurred_at, created_at, updated_at
		FROM transactions
		WHERE id = $1`
	var tx model.Transaction
	err := r.db.QueryRow(context.Background(), query, id).Scan(
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

func (r *PostgresTransactionRepository) UpdateTransaction(tx model.Transaction) (model.Transaction, error) {
	const query = `
		UPDATE transactions
		SET account_id = $2, amount = $3, currency = $4, category = $5, description = $6, occurred_at = $7, created_at = $8, updated_at = $9
		WHERE id = $1`
	result, err := r.db.Exec(context.Background(), query, tx.ID, tx.AccountID, tx.Amount, tx.Currency, tx.Category, tx.Description, tx.OccurredAt, tx.CreatedAt, tx.UpdatedAt)
	if err != nil {
		return model.Transaction{}, err
	}
	if result.RowsAffected() == 0 {
		return model.Transaction{}, storage.ErrNotFound
	}
	return tx, nil
}

func (r *PostgresTransactionRepository) DeleteTransaction(id string) error {
	const query = `DELETE FROM transactions WHERE id = $1`
	result, err := r.db.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return storage.ErrNotFound
	}
	return nil
}

func (r *PostgresTransactionRepository) ListTransactions() []model.Transaction {
	const query = `
		SELECT id, account_id, amount, currency, category, description, occurred_at, created_at, updated_at
		FROM transactions`
	rows, err := r.db.Query(context.Background(), query)
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

func (r *PostgresBudgetRepository) CreateBudget(budget model.Budget) (model.Budget, error) {
	const query = `
		INSERT INTO budgets (id, name, amount, currency, period, start_date, end_date, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err := r.db.Exec(context.Background(), query, budget.ID, budget.Name, budget.Amount, budget.Currency, budget.Period, budget.StartDate, budget.EndDate, budget.CreatedAt, budget.UpdatedAt)
	if err != nil {
		return model.Budget{}, err
	}
	return budget, nil
}

func (r *PostgresBudgetRepository) GetBudget(id string) (model.Budget, error) {
	const query = `
		SELECT id, name, amount, currency, period, start_date, end_date, created_at, updated_at
		FROM budgets
		WHERE id = $1`
	var budget model.Budget
	err := r.db.QueryRow(context.Background(), query, id).Scan(
		&budget.ID,
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

func (r *PostgresBudgetRepository) UpdateBudget(budget model.Budget) (model.Budget, error) {
	const query = `
		UPDATE budgets
		SET name = $2, amount = $3, currency = $4, period = $5, start_date = $6, end_date = $7, created_at = $8, updated_at = $9
		WHERE id = $1`
	result, err := r.db.Exec(context.Background(), query, budget.ID, budget.Name, budget.Amount, budget.Currency, budget.Period, budget.StartDate, budget.EndDate, budget.CreatedAt, budget.UpdatedAt)
	if err != nil {
		return model.Budget{}, err
	}
	if result.RowsAffected() == 0 {
		return model.Budget{}, storage.ErrNotFound
	}
	return budget, nil
}

func (r *PostgresBudgetRepository) DeleteBudget(id string) error {
	const query = `DELETE FROM budgets WHERE id = $1`
	result, err := r.db.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return storage.ErrNotFound
	}
	return nil
}

func (r *PostgresBudgetRepository) ListBudgets() []model.Budget {
	const query = `
		SELECT id, name, amount, currency, period, start_date, end_date, created_at, updated_at
		FROM budgets`
	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil
	}
	defer rows.Close()

	items := []model.Budget{}
	for rows.Next() {
		var budget model.Budget
		if err := rows.Scan(
			&budget.ID,
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
