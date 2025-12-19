COMPOSE ?= docker compose
POSTGRES_DSN ?= postgres://postgres:postgres@postgres:5432/ledger?sslmode=disable

.PHONY: up down start migrate migrate-ledger

up:
	$(COMPOSE) up -d

down:
	$(COMPOSE) down

start: up migrate

migrate: migrate-ledger

migrate-ledger:
	$(COMPOSE) run --rm ledger sh -c "go install github.com/pressly/goose/v3/cmd/goose@latest && POSTGRES_DSN=$(POSTGRES_DSN) /app/ledger/scripts/migrate.sh"
