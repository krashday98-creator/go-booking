# Booking API

Тестовое задание: REST API для комнат и бронирований

## Запуск

```bash
docker compose up --build
```

API: http://localhost:8080

Swagger UI: http://localhost:8080/swagger/index.html

## Swagger (Сгененировал с помощью Cursor, чтобы было легко тестировать)

Документация генерируется через [swaggo/swag](https://github.com/swaggo/swag).

```bash
go install github.com/swaggo/swag/cmd/swag@v1.16.4
swag init -g cmd/api/main.go -o docs --parseDependency --parseInternal
go run ./cmd/api
```

После изменения эндпоинтов или аннотаций перезапустите `swag init`.

## Авторизация

Администратор - `X-Admin-Key` - `admin-secret-key` 
Пользователь - `X-User-ID` - любой идентификатор пользователя 

## API


**Создание комнаты:**

Класс комнаты: `standard` или `deluxe`.

```json
{
  "class": "deluxe",
  "cost": 15000.00,
  "description": "Комната с видом на море"
}
```


**Бронирование:**

```json
{
  "room_id": "uuid-комнаты"
}
```

## Проверка (Windows)

Запуск: `docker compose up --build`. БД с хоста: `localhost:5433`, логин/пароль/БД — `booking`. Таблицы `rooms`, `bookings` появляются после первого старта API.

