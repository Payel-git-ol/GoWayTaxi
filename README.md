# GoWayTaxi

## Описание
GoWayTaxi — микросервисный бэкенд для такси‑сервиса, написанный на Go.  
Проект построен с использованием **Fiber**, **Kafka**, **Postgres**, **Prometheus/Grafana/Loki** и **Docker Compose**.  
Основная цель — обеспечить масштабируемую и надёжную архитектуру для регистрации водителей и пользователей, обработки заказов, расчёта стоимости и мониторинга состояния системы.

## Архитектура
Сервисы проекта:
- **Auth Service**: регистрация и авторизация пользователей (driver, rider), JWT‑токены.
- **User Service**: управление профилями пользователей.
- **Driver Service**: управление профилями водителей.
- **Order Service**: создание и обработка заказов.
- **Pricing Service**: расчёт стоимости поездки.
- **Rider Service**: логика взаимодействия с пассажирами.
- **Kafka**: брокер сообщений для асинхронного взаимодействия.
- **Postgres**: основная база данных для каждого сервиса.
- **Prometheus**: сбор метрик со всех сервисов.
- **Grafana**: визуализация метрик и логов.
- **Loki + Promtail**: сбор и хранение логов.

## Технологии
- **Go + Fiber** — быстрый HTTP‑фреймворк.
- **JWT** — авторизация и безопасность.
- **Kafka** — обмен событиями между сервисами.
- **Postgres** — хранение данных.
- **Prometheus/Grafana/Loki** — мониторинг и логирование.
- **Docker Compose** — оркестрация сервисов.

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

## Мониторинг
**Метрики сервисов доступны на порту 9100 каждого контейнера (/metrics)**.</br>
**Prometheus автоматически скрапит метрики по сервисам: auth-service, user-service, ride-service, pricing-service**.</br>
**Grafana подключена к Prometheus и Loki, готова к визуализации метрик и логов**.</br>
**Логи собираются Promtail и доступны в Grafana через источник Loki**.</br>