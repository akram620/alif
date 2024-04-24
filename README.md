# Инструкция по запуску проекта

## Запуск проекта через Docker

Для запуска проекта через Docker выполните следующие шаги:

1. В файле `.env` в корне проекта укажите переменные окружения:
   ```
   DB_URL="postgres://admin:Dskdhnjl**(0@psql_alif:5432/alif_task?sslmode=disable"
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
```bash
docker run -d --rm --name psql_alif \
-p 5432:5432 \
-e POSTGRES_USER=admin \
-e POSTGRES_PASSWORD="Dskdhnjl**(0" \
-e POSTGRES_DB=alif_task \
-v alif_postgresql_alif_data:/var/lib/postgresql/data \
postgres:16
```

Укажите переменные окружения в файле `.env` в корне проекта:
   ```
   DB_URL="postgres://admin:Dskdhnjl**(0@localhost:5432/alif_task?sslmode=disable"
   ```

#### Запуск базы данных без Docker

Для запуска базы данных без Docker выполните следующие шаги:

1. Установите PostgreSQL.
2. Создайте базу данных `alif_task`.
3. Создайте пользователя `admin`.
4. Укажите переменные окружения в файле `.env` в корне проекта:
   ```
   DB_URL="postgres://admin:Dskdhnjl**(0@localhost:5432/alif_task?sslmode=disable"
   ```