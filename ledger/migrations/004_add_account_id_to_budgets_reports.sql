-- +goose Up
ALTER TABLE budgets ADD COLUMN IF NOT EXISTS account_id TEXT NOT NULL;
ALTER TABLE reports ADD COLUMN IF NOT EXISTS account_id TEXT NOT NULL;

CREATE INDEX IF NOT EXISTS budgets_account_id_idx ON budgets (account_id);
CREATE INDEX IF NOT EXISTS reports_account_id_idx ON reports (account_id);

-- +goose Down
DROP INDEX IF EXISTS reports_account_id_idx;
DROP INDEX IF EXISTS budgets_account_id_idx;
ALTER TABLE reports DROP COLUMN IF EXISTS account_id;
ALTER TABLE budgets DROP COLUMN IF EXISTS account_id;
