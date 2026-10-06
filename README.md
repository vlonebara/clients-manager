# Clients Manager

Небольшое HTTP-приложение на Go для управления пользователями. Данные пользователей хранятся в SQLite.

## Стек

- Go 1.26.1
- SQLite
- `database/sql`
- [chi](https://github.com/go-chi/chi) — HTTP-роутер
- [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) — безопасное хеширование паролей

## Структура проекта

```text
clients-manager/
├── cmd/
│   └── app/
│       └── main.go                  # Точка входа
├── db/
│   └── app                         # Файл SQLite-базы
├── internal/
│   ├── handlers/
│   │   ├── AddUserHandler.go       # Создание пользователя
│   │   ├── GetUsersHandler.go      # Получение списка пользователей
│   │   └── TestHandler.go          # Проверка работы сервера
│   └── server/
│       └── server.go               # Настройка и запуск сервера
├── go.mod
└── go.sum
```

## Запуск

Клонируйте репозиторий и перейдите в его корневой каталог:

```bash
git clone https://github.com/vlonebara/clients-manager.git
cd clients-manager
```

Установите зависимости и запустите приложение:

```bash
go mod download
go run ./cmd/app
```

После запуска сервер доступен по адресу:

```text
http://localhost:8080
```

## Маршруты

### Проверка работы сервера

```http
GET /test
```

Маршрут возвращает host из входящего HTTP-запроса.

Пример:

```bash
curl http://localhost:8080/test
```

Пример ответа:

```text
localhost:8080
```

### Создание пользователя

```http
POST /addUser
Content-Type: application/json
```

Пример тела запроса:

```json
{
  "name": "Кирилл",
  "email": "kirill@example.com",
  "role": "admin",
  "login": "kirill",
  "password": "123456789qwer"
}
```

Пример запроса через `curl`:

```bash
curl -X POST http://localhost:8080/addUser \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Кирилл",
    "email": "kirill@example.com",
    "role": "admin",
    "login": "kirill",
    "password": "123456789qwer"
  }'
```

При успешном создании сервер возвращает статус `201 Created` и JSON с данными нового пользователя. Пароль не возвращается в ответе: в базе сохраняется только bcrypt-хеш.

Пример ответа:

```json
{
  "id": 1,
  "name": "Кирилл",
  "email": "kirill@example.com",
  "role": "admin",
  "login": "kirill"
}
```

Допустимые роли:

- `manager`
- `admin`

### Получение списка пользователей

```http
GET /users
```

Маршрут возвращает список пользователей в JSON. Пароль и `password_hash` в ответ не включаются. Пользователи возвращаются в порядке возрастания `id`.

Пример:

```bash
curl http://localhost:8080/users
```

Пример ответа:

```json
[
  {
    "id": 1,
    "name": "Кирилл",
    "email": "kirill@example.com",
    "role": "admin",
    "login": "kirill"
  }
]
```

## Таблица пользователей

Для работы создания и получения пользователей в SQLite должна быть таблица `users`:

```sql
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    role TEXT NOT NULL CHECK (role IN ('manager', 'admin')),
    login TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL
);
```

## Важное замечание

При отправке данных используй `POST` и JSON в теле запроса. Не передавай пароль в URL-параметрах, например `?password=...`.
