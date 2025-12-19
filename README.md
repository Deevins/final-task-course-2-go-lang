# Gateway API

HTTP gateway для сервисов авторизации и Ledger.

Базовый адрес: `http://localhost:8081`

## Запуск через Docker Compose

```bash
docker compose up
```

## Запуск через Make

```bash
make start
```

Команды:

- `make up` — поднять все сервисы через Docker Compose
- `make migrate` — накатить миграции (требуются для Ledger)
- `make start` — поднять сервисы и выполнить миграции
- `make down` — остановить сервисы

После запуска сервисы доступны по адресам:

- Gateway HTTP: `http://localhost:8081`
- Auth HTTP: `http://localhost:8082`
- Auth gRPC: `localhost:9092`
- Ledger HTTP: `http://localhost:8083`
- Ledger gRPC: `localhost:9091`
- Postgres: `localhost:5432` (DB `ledger`, пользователь `postgres`, пароль `postgres`)
- Redis: `localhost:6379`

## Миграции (Auth + Ledger)

Запуск миграций Auth через Docker Compose:

```bash
docker compose run --rm auth sh -c "go install github.com/pressly/goose/v3/cmd/goose@latest && goose -dir /app/auth/migrations postgres \"postgres://postgres:postgres@postgres:5432/ledger?sslmode=disable\" up"
```

Запуск миграций Ledger через Docker Compose:

```bash
docker compose run --rm ledger sh -c "go install github.com/pressly/goose/v3/cmd/goose@latest && POSTGRES_DSN=postgres://postgres:postgres@postgres:5432/ledger?sslmode=disable /app/ledger/scripts/migrate.sh"
```

## Авторизация (JWT)

Gateway работает с JWT: после `POST /api/auth/signin` вы получаете токен и используете
его для всех маршрутов Ledger через заголовок `Authorization: Bearer <jwt>`.

```bash
curl -H "Authorization: Bearer <jwt>" \
  http://localhost:8081/api/ledger/transactions
```

## Маршруты и примеры запросов

### Auth

- `POST /api/auth/signup` — регистрация пользователя.
- `POST /api/auth/signin` — аутентификация и получение JWT.

Пример регистрации:

```bash
curl -X POST http://localhost:8081/api/auth/signup \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "secret",
    "name": "Иван Иванов"
  }'
```

Пример получения JWT:

```bash
curl -X POST http://localhost:8081/api/auth/signin \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "secret"
  }'
```

### Ledger (требует Bearer JWT)

Требуется заголовок `Authorization: Bearer <jwt>` для всех маршрутов ниже.

- Транзакции:
  - `GET /api/ledger/transactions`
  - `POST /api/ledger/transactions`
  - `GET /api/ledger/transactions/{id}`
  - `PUT /api/ledger/transactions/{id}`
  - `PATCH /api/ledger/transactions/{id}`
  - `DELETE /api/ledger/transactions/{id}`
- Бюджеты:
  - `GET /api/ledger/budgets`
  - `POST /api/ledger/budgets`
  - `GET /api/ledger/budgets/{id}`
  - `PUT /api/ledger/budgets/{id}`
  - `PATCH /api/ledger/budgets/{id}`
  - `DELETE /api/ledger/budgets/{id}`
- Отчеты:
  - `GET /api/ledger/reports`
  - `POST /api/ledger/reports`
  - `GET /api/ledger/reports/{id}`
  - `PUT /api/ledger/reports/{id}`
  - `PATCH /api/ledger/reports/{id}`
  - `DELETE /api/ledger/reports/{id}`
- Импорт/экспорт:
  - `POST /api/ledger/import`
  - `GET /api/ledger/export`

Примечание: интеграции с GSheets нет — используется только CSV импорт/экспорт.

Пример списка транзакций:

```bash
curl -H "Authorization: Bearer <jwt>" \
  http://localhost:8081/api/ledger/transactions
```

Пример создания транзакции (если `account_id` не передан, берется из JWT):

```bash
curl -X POST http://localhost:8081/api/ledger/transactions \
  -H "Authorization: Bearer <jwt>" \
  -H "Content-Type: application/json" \
  -d '{
    "amount": 1250.5,
    "currency": "RUB",
    "category": "Продукты",
    "description": "Покупка в магазине",
    "occurred_at": "2024-01-01T10:00:00Z"
  }'
```
