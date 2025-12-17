# GoWayTaxi

## Описание
GoWayTaxi — микросервисный бэкенд для такси‑сервиса, написанный на Go.  
Проект построен с использованием Fiber, Kafka, Postgres и Docker Compose.  
Основная цель — обеспечить масштабируемую и надёжную архитектуру для регистрации водителей и пользователей, обработки заказов и взаимодействия между сервисами.

## Архитектура
Сервисы проекта:
- Auth Service: регистрация и авторизация пользователей (driver, rider), JWT‑токены.
- User Service: управление профилями пользователей.
- Driver Service: управление профилями водителей.
- Order Service: создание и обработка заказов.
- Pricing Service: расчёт стоимости поездки.
- Rider Service: логика взаимодействия с пассажирами.
- Kafka: брокер сообщений для асинхронного взаимодействия.
- Postgres: основная база данных.

## Технологии
- Go + Fiber — быстрый HTTP‑фреймворк.
- JWT — авторизация и безопасность.
- Kafka — обмен событиями между сервисами.
- Postgres — хранение данных.
- Docker Compose — оркестрация сервисов.

## Запуск проекта
1. Клонировать репозиторий:
   ```bash
   git clone https://github.com/Payel-git-ol/GoWayTaxi.git
   cd GoWayTaxi

2. Запустить сервисы через Docker Compose:

   ```bash
   docker-compose up --build
   ```
3. После запуска будут доступны:

   Auth API: http://localhost:8080/auth
   
   User API: http://localhost:8081/user
   
   Driver API: http://localhost:8082/driver
   
   Order API: http://localhost:8083/order

## Примеры API
   Регистрация пользователя
   ```
   POST /auth/register
   Content-Type: application/json
   ```

   ```json
  {
      "username": "pavel",
      "password": "securepass",
      "role": "user"
  }
   ```

   Регистрация водителя
   ```
   POST /auth/register 
   Content-Type: application/json
   ```

   ```json
   {
      "username": "driver1", 
      "password": "securepass", 
      "role": "driver"
  }
   ```

## Авторизация

```
POST /auth/login
Content-Type: application/json 
```

```json
{ 
   "username": "pavel",
   "password": "securepass"
}
```

Ответ:
```json
{
  "token": "jwt-token-here"
}
```