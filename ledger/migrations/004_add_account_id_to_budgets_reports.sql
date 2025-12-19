-- +goose Up
ALTER TABLE budgets ADD COLUMN account_id TEXT NOT NULL;
ALTER TABLE reports ADD COLUMN account_id TEXT NOT NULL;

-- +goose Down
ALTER TABLE reports DROP COLUMN account_id;
ALTER TABLE budgets DROP COLUMN account_id;
