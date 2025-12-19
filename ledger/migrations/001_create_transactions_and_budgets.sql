-- +goose Up
CREATE TABLE IF NOT EXISTS transactions (
    id TEXT PRIMARY KEY,
    account_id TEXT NOT NULL,
    amount DOUBLE PRECISION NOT NULL,
    currency TEXT NOT NULL,
    category TEXT NOT NULL,
    description TEXT NOT NULL,
    occurred_at TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL
);

CREATE INDEX IF NOT EXISTS transactions_account_id_idx ON transactions (account_id);
CREATE INDEX IF NOT EXISTS transactions_category_idx ON transactions (category);
CREATE INDEX IF NOT EXISTS transactions_occurred_at_idx ON transactions (occurred_at);

CREATE TABLE IF NOT EXISTS budgets (
    id TEXT PRIMARY KEY,
    account_id TEXT NOT NULL,
    name TEXT NOT NULL,
    amount DOUBLE PRECISION NOT NULL,
    currency TEXT NOT NULL,
    period TEXT NOT NULL,
    start_date TIMESTAMPTZ NOT NULL,
    end_date TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL
);

CREATE INDEX IF NOT EXISTS budgets_account_id_idx ON budgets (account_id);
CREATE INDEX IF NOT EXISTS budgets_name_idx ON budgets (name);
CREATE INDEX IF NOT EXISTS budgets_currency_idx ON budgets (currency);
CREATE INDEX IF NOT EXISTS budgets_period_idx ON budgets (period);

-- +goose Down
DROP TABLE IF EXISTS transactions;
DROP TABLE IF EXISTS budgets;

