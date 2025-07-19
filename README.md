# **🐾 Go Pet Shop - Документация**  

**Go Pet Shop** — это веб-приложение для управления интернет-магазином товаров для домашних животных 🐶🐱🦜.  
Проект написан на **Go** с использованием **Gin** (HTTP-фреймворк) и **PostgreSQL** (база данных).  

---

## **🚀 Основные возможности**  

### **👥 Управление пользователями**  
- `POST /users` — Создать нового пользователя.  
- `GET /users` — Получить список всех пользователей.  

### **🛍️ Управление товарами**  
- `POST /products` — Добавить новый товар.  
- `GET /products` — Получить список всех товаров.  
- `PUT /products/:id` — Обновить информацию о товаре.  
- `DELETE /products/:id` — Удалить товар.  

### **📦 Управление заказами**  
- `POST /orders` — Создать заказ.  
- `GET /orders` — Получить заказы пользователя.  
- `GET /orders/history` — Получить историю заказов.  

### **💳 Оформление заказа**  
- `POST /checkout` — Оформить заказ (проверка наличия товара + транзакция).  

### **📊 Аналитика**  
- `GET /analytics/popular-products` — Топ популярных товаров.  

### **💸 Платежи**  
- `POST /payments` — Обработать платеж.  

### **🩺 Статус сервера**  
- `GET /status` — Проверить работоспособность сервера.  

---

## **🛠 Технологии**  
| Компонент       | Технология       |
|----------------|-----------------|
| **Backend**    | Go (Gin)        |
| **База данных** | PostgreSQL      |
| **Миграции**   | golang-migrate  |
| **Тестирование** | Testify        |
| **Конфигурация** | YAML           |

---

## **⚙️ Установка и запуск**  

### **1️⃣ Клонирование репозитория**  
```bash
git clone https://github.com/<ваш-репозиторий>/go_pet_shop.git
cd go_pet_shop
```

### **2️⃣ Настройка базы данных**  
1. Установите **PostgreSQL** и создайте БД:  
   ```sql
   CREATE DATABASE petshop;
   ```
2. Настройте подключение в `local.yaml`:  
   ```yaml
   db:
     url: "postgres://user:password@localhost:5432/petshop?sslmode=disable"
   ```

### **3️⃣ Применение миграций**  
```bash
go run cmd/migrator/main.go --migrations-path "./migrations"
```

### **4️⃣ Запуск сервера**  
```bash
go run cmd/app/main.go
```

### **5️⃣ Тестирование**  
```bash
go test ./tests/...
```

---

## **📡 API-эндпоинты**  

### **👤 Пользователи**  
```bash
# Создать пользователя
curl -X POST http://localhost:8080/users -d '{"name": "Alice", "email": "alice@example.com"}'

# Получить всех пользователей
curl -X GET http://localhost:8080/users
```

### **🛒 Товары**  
```bash
# Добавить товар
curl -X POST http://localhost:8080/products -d '{"name": "Корм для кошек", "price": 599, "stock": 50}'

# Получить список товаров
curl -X GET http://localhost:8080/products
```

### **📦 Заказы**  
```bash
# Оформить заказ
curl -X POST http://localhost:8080/checkout -d '{
  "user_email": "alice@example.com",
  "items": [{"product_id": 1, "quantity": 2}]
}'
```

---

## **📂 Структура проекта**  
```
├── cmd/            # Точки входа
│   ├── app/        # Основной сервер
│   └── migrator/   # Миграции БД
├── config/         # Конфигурация
├── internal/       # Бизнес-логика
│   ├── handlers/   # HTTP-обработчики
│   └── storage/    # Работа с БД
├── models/         # Модели данных
├── migrations/     # SQL-миграции
└── tests/          # Тесты
```

---

## **📜 Примеры запросов**  

### **1. Добавление товара**  
```bash
curl -X POST http://localhost:8080/products -H "Content-Type: application/json" -d '{
  "name": "Игрушка для собак",
  "price": 299,
  "stock": 100
}'
```

### **2. Оформление заказа**  
```bash
curl -X POST http://localhost:8080/checkout -H "Content-Type: application/json" -d '{
  "user_email": "alice@example.com",
  "items": [
    {"product_id": 1, "quantity": 2},
    {"product_id": 2, "quantity": 1}
  ]
}'
```

