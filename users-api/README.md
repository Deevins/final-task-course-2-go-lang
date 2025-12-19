## Users API (Gin)

Минималистичный микросервис на Go:
- CRUD по пользователям (Gin + чистая архитектура, in-memory хранилище).

### Структура

```
cmd/
  server/
    main.go           # точка входа
internal/
  handler/           # HTTP-слой (Gin)
  service/           # бизнес-логика
  repository/        # репозиторий (in-memory)
  model/             # доменные модели и DTO
  docs/              # минимальные Swagger-доки (можно перегенерировать swag CLI)
```

### Установка и запуск

Установка зависимостей и генерация Swagger (опционально):

```bash
make deps tidy
make swag
```

Запуск сервера:

```bash
make run
# сервер слушает :8080
```

Документация:

- Swagger UI: `http://localhost:8080/swagger/index.html`
- ReDoc: `http://localhost:8080/redoc`

### User CRUD

- GET    `/api/v1/users` — список
- POST   `/api/v1/users` — создать
- GET    `/api/v1/users/{id}` — получить по ID
- PATCH  `/api/v1/users/{id}` — частичное обновление
- DELETE `/api/v1/users/{id}` — удалить

Примеры запросов:

```bash
# создать
curl -s -X POST http://localhost:8080/api/v1/users \
  -H 'Content-Type: application/json' \
  -d '{"username":"alice","email":"alice@example.com"}'

# список
curl -s http://localhost:8080/api/v1/users | jq .

# получить по id
curl -s http://localhost:8080/api/v1/users/<id>

# частичное обновление
curl -s -X PATCH http://localhost:8080/api/v1/users/<id> \
  -H 'Content-Type: application/json' \
  -d '{"username":"alice2"}'

# удалить
curl -s -X DELETE http://localhost:8080/api/v1/users/<id> -i
```

- Swagger аннотации находятся в `internal/handler/*.go`.
- Перегенерация OpenAPI: `swag init -g cmd/server/main.go -o internal/docs`.
