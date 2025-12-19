COMPOSE ?= docker compose
AUTH_POSTGRES_DSN ?= postgres://postgres:postgres@auth-postgres:5432/auth?sslmode=disable
LEDGER_POSTGRES_DSN ?= postgres://postgres:postgres@ledger-postgres:5432/ledger?sslmode=disable

.PHONY: up down start migrate migrate-auth migrate-ledger migrate-db

up:
	$(COMPOSE) up -d

down:
	$(COMPOSE) down

start: up migrate

migrate: migrate-db migrate-auth migrate-ledger

migrate-db:
	$(COMPOSE) up -d auth-postgres ledger-postgres

migrate-auth:
	$(COMPOSE) run --rm auth sh -c "go install github.com/pressly/goose/v3/cmd/goose@latest && goose -dir /app/auth/migrations postgres \"$(AUTH_POSTGRES_DSN)\" up"

migrate-ledger:
	$(COMPOSE) run --rm ledger sh -c "go install github.com/pressly/goose/v3/cmd/goose@latest && POSTGRES_DSN=$(LEDGER_POSTGRES_DSN) /app/ledger/scripts/migrate.sh"
