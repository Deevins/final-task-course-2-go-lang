-- +goose Up
CREATE TABLE IF NOT EXISTS reports (
    id TEXT PRIMARY KEY,
    account_id TEXT NOT NULL,
    name TEXT NOT NULL,
    period TEXT NOT NULL,
    generated_at TIMESTAMPTZ NOT NULL,
    total_income DOUBLE PRECISION NOT NULL,
    total_expense DOUBLE PRECISION NOT NULL,
    currency TEXT NOT NULL,
    categories JSONB NOT NULL DEFAULT '[]'
);

CREATE INDEX IF NOT EXISTS reports_account_id_idx ON reports (account_id);
CREATE INDEX IF NOT EXISTS reports_period_idx ON reports (period);
CREATE INDEX IF NOT EXISTS reports_currency_idx ON reports (currency);
CREATE INDEX IF NOT EXISTS reports_generated_at_idx ON reports (generated_at);

-- +goose Down
DROP TABLE IF EXISTS reports;
