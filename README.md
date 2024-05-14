# Инструкция по запуску проекта

## Запуск проекта через Docker

Для запуска проекта через Docker выполните следующие шаги:

1. В файле `.env` укажите переменные окружения:
   ```
   DB_URL="postgres://admin:Dskdhnjl**(0@work_scheduler_pg:5432/work_scheduler?sslmode=disable"
   ```
2. Выполните команду:
   ```
   docker-compose up --build
   ```
   в корне проекта.

## Запуск проекта без Docker

### Запуск базы данных

#### Запуск базы данных через Docker

Для запуска базы данных через Docker выполните следующие действия:
1. Выполните команду:
```bash
docker run -d --rm --name work_scheduler_pg \
-p 5432:5432 \
-e POSTGRES_USER=admin \
-e POSTGRES_PASSWORD="Dskdhnjl**(0" \
-e POSTGRES_DB=work_scheduler \
-v work_scheduler_data:/var/lib/postgresql/data \
postgres:16
```

2. Укажите переменные окружения в файле `.env`:
   ```
   DB_URL="postgres://admin:Dskdhnjl**(0@localhost:5432/work_scheduler?sslmode=disable"
   ```

#### Запуск базы данных без Docker

Для запуска базы данных без Docker выполните следующие шаги:

1. Установите PostgreSQL.
2. Создайте базу данных `work_scheduler`.
3. Создайте пользователя `admin`.
4. Укажите переменные окружения в файле `.env`:
   ```
   DB_URL="postgres://admin:Dskdhnjl**(0@localhost:5432/work_scheduler?sslmode=disable"
   ```

---
Выполните команду:
```bash
go mod download
```
```bash
go run cmd/server/main.go
```
или 
```bash
sh run.sh
```
