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
  - Бюджет задается только на месяц (поле `month` — дата первого дня месяца в формате RFC3339), а категория определяется полем `name`.
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

Период отчета (`period`) поддерживает оба формата:

- `YYYY-MM` (месяц, например `2024-01`)
- `start/end` — две даты, разделенные `/`, в формате `YYYY-MM-DD` или RFC3339
  (например `2024-01-01/2024-01-31` или `2024-01-01T00:00:00Z/2024-01-31T23:59:59Z`)

Пример создания отчета с месячным периодом:

```bash
curl -X POST http://localhost:8081/api/ledger/reports \
  -H "Authorization: Bearer <jwt>" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Январь 2024",
    "period": "2024-01",
    "generated_at": "2024-01-31T23:59:59Z",
    "currency": "RUB"
  }'
```

Пример создания отчета с диапазоном дат:

```bash
curl -X POST http://localhost:8081/api/ledger/reports \
  -H "Authorization: Bearer <jwt>" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Январь 2024",
    "period": "2024-01-01/2024-01-31",
    "generated_at": "2024-01-31T23:59:59Z",
    "currency": "RUB"
  }'
```

Пример ответа отчета с разбивкой по категориям (при `budget_amount = 0` поле
`budget_usage_percent` возвращается как `null`):

Правило расчета `budget_amount` в отчете: для каждого месячного бюджета берется
пересечение дат отчета и месяца бюджета по календарным дням (границы включаются),
доля вычисляется как `пересекающиеся_дни / дни_в_месяце`, затем доли по месяцам
суммируются.

```json
{
  "id": "11111111-1111-1111-1111-111111111111",
  "account_id": "22222222-2222-2222-2222-222222222222",
  "name": "Январь 2024",
  "period": "2024-01",
  "generated_at": "2024-01-31T23:59:59Z",
  "total_income": 50000,
  "total_expense": 30000,
  "currency": "RUB",
  "categories": [
    {
      "category": "Продукты",
      "total_expense": 30000,
      "budget_amount": 0,
      "budget_usage_percent": null
    }
  ]
}
```
