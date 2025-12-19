-- +goose Up
ALTER TABLE budgets ADD COLUMN account_id TEXT NOT NULL;
ALTER TABLE reports ADD COLUMN account_id TEXT NOT NULL;

CREATE INDEX IF NOT EXISTS budgets_account_id_idx ON budgets (account_id);
CREATE INDEX IF NOT EXISTS reports_account_id_idx ON reports (account_id);

-- +goose Down
DROP INDEX IF EXISTS reports_account_id_idx;
DROP INDEX IF EXISTS budgets_account_id_idx;
ALTER TABLE reports DROP COLUMN account_id;
ALTER TABLE budgets DROP COLUMN account_id;
