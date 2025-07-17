# 🧩 Social Network Microservices — Go + Docker + PostgreSQL

Микросервисный проект "Социальная сеть", реализованный на Go с использованием `Gin`, `Docker` и `PostgreSQL`.

## 📌 Функциональность

- 👤 Регистрация пользователей
- 🔐 Вход в систему
- 📝 Создание постов
- 📄 Просмотр всех постов
- 🔍 Просмотр постов конкретного пользователя
- ❤️ Лайк поста
- 💬 Комментарий к посту

## 🛠️ Технологии

- Go 1.22+
- Gin (HTTP-фреймворк)
- PostgreSQL 14
- Docker + Docker Compose

## 🧱 Архитектура

Проект состоит из трёх микросервисов:

```
api-gateway     → принимает HTTP-запросы от клиента
    |
    ↓
processor       → бизнес-логика, маршрутизация
    |
    ↓
db-service      → работа с PostgreSQL
    |
    ↓
PostgreSQL      → хранение пользователей, постов, лайков и комментариев
```

## 📁 Структура проекта

```
social-network/
│
├── api-gateway/        # Приём HTTP-запросов
│   ├── main.go
│   ├── handlers/
│   ├── routes/
│   └── Dockerfile
│
├── processor/          # Бизнес-логика
│   ├── main.go
│   ├── handlers/
│   ├── utils/
│   ├── routes/
│   └── Dockerfile
│
├── db-service/         # Работа с БД
│   ├── main.go
│   ├── database/
│   ├── routes/
│   ├── models/
│   ├── handlers/
│   └── Dockerfile
│
├── docker-compose.yml  # Поднятие всех сервисов
└── init.sql            # Скрипт инициализации базы данных
```

## 🚀 Запуск проекта

### 1. Клонируй репозиторий

```bash
git clone https://github.com/Gayana5/social-network.git
cd social-network
```

### 2. Запусти проект с Docker Compose

```bash
docker-compose up --build
```

Проект поднимет все микросервисы: `api-gateway`, `processor`, `db-service`, `postgres`.

## 🧪 Тестирование API

Можно использовать Postman, Insomnia или `curl`.

### 🔐 Регистрация

```
POST http://localhost:8080/signUp
Content-Type: application/json
```

```json
{
  "username": "user1",
  "password": "pass123"
}
```

### 🔐 Вход

```
POST http://localhost:8080/signIn
Content-Type: application/json
```

```json
{
  "username": "user1",
  "password": "pass123"
}
```

### 📝 Создание поста

```
POST http://localhost:8080/post
Content-Type: application/json
```

```json
{
  "user_id": 1,
  "content": "Текст поста"
}
```

### 📄 Получить все посты

```
GET http://localhost:8080/posts
```

### 🔍 Посты пользователя

```
GET http://localhost:8080/posts/1
```

### ❤️ Лайк поста

```
POST http://localhost:8080/post/1/like
```

### 💬 Комментарий к посту

```
POST http://localhost:8080/post/1/comment
Content-Type: application/json
```

```json
{
  "user_id": 1,
  "content": "Текст комментария"
}
```

## 🐳 Docker-команды

Остановить и удалить контейнеры:

```bash
docker-compose down
```

Полная пересборка и запуск:

```bash
docker-compose down -v
docker-compose up --build
```
