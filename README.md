# Booking

Запуск
1. Поднять базу данных:

```bash
docker compose up -d db
```

2. Запустить сервер:

```bash
go run ./cmd/api
```

Сервер: `http://localhost:8080`

Эндпоинты

Получить все комнаты

`GET /rooms`

Получить комнату по id

`GET /rooms/:id`

Если не найдена: `404 Not Found`

```json
{"error": "room not found"}
```

Создать комнату

`POST /rooms`

Удалить комнату

`DELETE /rooms/:id`

Забронировать комнату

`POST /bookings`

Если комната занята:

`409 Conflict`
Если комната не найдена:

`404 Not Found`

Получить информацию о броинровании

`GET /bookings?user_id=1`

Получить информацию о бронировании
GET /bookings/:id

Получить все брони пользователя
GET /bookings?user_id=1
