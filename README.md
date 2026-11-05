# Clients Manager

Clients Manager — небольшой HTTP-сервис на Go для создания пользователей и просмотра списка. Данные хранятся в SQLite; пароли перед сохранением хешируются с помощью bcrypt.

## Возможности

- создать пользователя через `POST /addUser`;
- получить список пользователей через `GET /users`;
- проверить доступность сервера через `GET /test`.

Пароль и его хеш не включаются в ответы API.

## Архитектура

```text
cmd/app/main.go                 точка входа и запуск сервера
internal/server/server.go       подключение SQLite и настройка HTTP-маршрутов
internal/handlers/              обработчики HTTP-запросов
  add_user.go                   создание пользователя и хеширование пароля
  get_users.go                  выдача списка пользователей
  test.go                       проверочный маршрут
db/app                         SQLite-база данных
```

Точка входа запускает сервер. Сервер открывает SQLite через `database/sql`, создаёт маршрутизатор chi и связывает маршруты с обработчиками. Обработчик создания пользователя преобразует JSON-запрос, хеширует пароль через bcrypt и записывает пользователя в базу. При чтении списка выбираются только публичные поля пользователя.

## Стек

- Go 1.26.1+
- `net/http` и `database/sql`
- [chi](https://github.com/go-chi/chi) для маршрутизации
- SQLite через `github.com/mattn/go-sqlite3`
- bcrypt через `golang.org/x/crypto/bcrypt`

## Запуск

Нужен Go 1.26.1 или новее. Выполняйте команды из корня репозитория: приложение открывает базу по относительному пути `./db/app`.

```bash
git clone https://github.com/vlonebara/clients-manager.git
cd clients-manager
go mod download
go run ./cmd/app
```

Сервер запустится на `http://localhost:8080`.

## API

### Проверка сервера

```bash
curl -i http://localhost:8080/test
```

Возвращает host запроса обычным текстом.

### Создание пользователя

```bash
curl -i -X POST http://localhost:8080/addUser \
  -H 'Content-Type: application/json' \
  -d '{"name":"Кирилл","email":"kirill@example.com","role":"manager","login":"kirill","password":"change-me"}'
```

Допустимые роли: `manager` и `admin`. При успешном создании возвращается `201 Created` и объект с полями `id`, `name`, `email`, `role` и `login`. Пароль в ответе отсутствует.

### Список пользователей

```bash
curl -i http://localhost:8080/users
```

Возвращает JSON-массив пользователей, отсортированный по `id`. В ответе нет пароля и его хеша.

## Схема SQLite

В базе `db/app` должны быть таблицы `users` и `clients`. Таблица `users` содержит `id`, `name`, `email`, `role`, `login` и `password_hash`; `email` и `login` уникальны, а `role` ограничена значениями `manager` и `admin`. Таблица `clients` сейчас не используется HTTP-маршрутами.
