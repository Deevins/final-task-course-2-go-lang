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

## Авторизация

Используйте Bearer JWT токен в заголовке `Authorization` для защищенных маршрутов Ledger.

```bash
curl -H "Authorization: Bearer <jwt>" \
  http://localhost:8081/api/ledger/transactions
```

## Маршруты и примеры запросов

### Auth

#### POST `/api/auth/signup`
Регистрация пользователя.

```bash
curl -X POST http://localhost:8081/api/auth/signup \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "secret",
    "name": "Иван Иванов"
  }'
```

#### POST `/api/auth/signin`
Аутентификация пользователя и получение JWT.

```bash
curl -X POST http://localhost:8081/api/auth/signin \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "secret"
  }'
```

### Ledger (требует Bearer JWT)

#### GET `/api/ledger/transactions`
Список транзакций пользователя.

```bash
curl -H "Authorization: Bearer <jwt>" \
  http://localhost:8081/api/ledger/transactions
```

#### POST `/api/ledger/transactions`
Создание транзакции (если `account_id` не передан, берется из JWT).

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

#### GET `/api/ledger/budgets`
Список бюджетов.

```bash
curl -H "Authorization: Bearer <jwt>" \
  http://localhost:8081/api/ledger/budgets
```

#### POST `/api/ledger/budgets`
Создание бюджета.

```bash
curl -X POST http://localhost:8081/api/ledger/budgets \
  -H "Authorization: Bearer <jwt>" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Еда",
    "amount": 10000,
    "currency": "RUB",
    "period": "monthly",
    "start_date": "2024-01-01T00:00:00Z",
    "end_date": "2024-01-31T23:59:59Z"
  }'
```

#### GET `/api/ledger/reports`
Список отчетов.

```bash
curl -H "Authorization: Bearer <jwt>" \
  http://localhost:8081/api/ledger/reports
```

#### POST `/api/ledger/reports`
Создание отчета.

```bash
curl -X POST http://localhost:8081/api/ledger/reports \
  -H "Authorization: Bearer <jwt>" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Январь 2024",
    "period": "2024-01/2024-01-31",
    "generated_at": "2024-01-31T23:59:59Z",
    "currency": "RUB"
  }'
```

#### POST `/api/ledger/import`
Импорт транзакций из CSV.

```bash
curl -X POST http://localhost:8081/api/ledger/import \
  -H "Authorization: Bearer <jwt>" \
  -H "Content-Type: application/json" \
  -d '{
    "csv_content": "account_id,amount,currency,category,description,occurred_at\n2222,1000,RUB,Продукты,Покупка,2024-01-01T10:00:00Z",
    "has_header": true
  }'
```

#### GET `/api/ledger/export`
Экспорт транзакций в CSV.

```bash
curl -H "Authorization: Bearer <jwt>" \
  http://localhost:8081/api/ledger/export
```
