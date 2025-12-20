COMPOSE ?= docker compose
AUTH_POSTGRES_DSN ?= postgres://postgres:postgres@auth-postgres:5432/auth?sslmode=disable
LEDGER_POSTGRES_DSN ?= postgres://postgres:postgres@ledger-postgres:5432/ledger?sslmode=disable

.PHONY: up down start migrate migrate-auth migrate-ledger migrate-db wait-ledger-db

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

migrate-ledger: wait-ledger-db
	$(COMPOSE) run --rm ledger sh -c "go install github.com/pressly/goose/v3/cmd/goose@latest && tr -d '\r' < /app/ledger/scripts/migrate.sh > /tmp/migrate.sh && chmod +x /tmp/migrate.sh && POSTGRES_DSN=$(LEDGER_POSTGRES_DSN) /tmp/migrate.sh"


wait-ledger-db:
	$(COMPOSE) exec -T ledger-postgres sh -c "until pg_isready -U postgres -d ledger >/dev/null 2>&1; do sleep 1; done"
