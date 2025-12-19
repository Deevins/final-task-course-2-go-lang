-- +goose Up
ALTER TABLE budgets ADD COLUMN IF NOT EXISTS month TIMESTAMPTZ;

ALTER TABLE budgets ALTER COLUMN month SET NOT NULL;
ALTER TABLE budgets DROP COLUMN IF EXISTS start_date;
ALTER TABLE budgets DROP COLUMN IF EXISTS end_date;

CREATE INDEX IF NOT EXISTS budgets_month_idx ON budgets (month);
CREATE UNIQUE INDEX IF NOT EXISTS budgets_account_name_month_idx ON budgets (account_id, name, month);

-- +goose Down
ALTER TABLE budgets ADD COLUMN IF NOT EXISTS start_date TIMESTAMPTZ;
ALTER TABLE budgets ADD COLUMN IF NOT EXISTS end_date TIMESTAMPTZ;

UPDATE budgets
SET start_date = date_trunc('month', month),
    end_date = (date_trunc('month', month) + interval '1 month' - interval '1 nanosecond')
WHERE start_date IS NULL;

ALTER TABLE budgets ALTER COLUMN start_date SET NOT NULL;
ALTER TABLE budgets ALTER COLUMN end_date SET NOT NULL;

ALTER TABLE budgets DROP COLUMN IF EXISTS month;
DROP INDEX IF EXISTS budgets_month_idx;
DROP INDEX IF EXISTS budgets_account_name_month_idx;
