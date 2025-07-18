## 🚀 Go Pet Shop V1

### 🗃️ **База данных PostgreSQL**
- ✅ Полностью настроена схема БД для интернет-магазина
- ✅ Реализованы все основные таблицы:
  - `products` (товары)
  - `users` (пользователи)
  - `orders` (заказы) 
  - `order_items` (позиции заказов)
  - `transactions` (платежи)
- ✅ Настроены миграции через `golang-migrate`
 
### ⚙️ **Ядро приложения** 
- ✅ Конфигурация через `local.yaml` 
- ✅ Подключение к БД через `sqlx` 
- ✅ Базовые модели данных (structs)

### 🛠️ **Storage-слой**
- ✅ Полностью реализован `ProductRepository`:
  - `CreateProduct()`  
  - `GetProductByID()`
  - `GetAllProducts()`
  - `UpdateProduct()`
  - `DeleteProduct()`
- ✅ Транзакционный метод `PlaceOrder()`:
  - Создание заказа
  - Проверка остатков
  - Обновление склада
  - Создание транзакции

### 📦 **Готовые зависимости**
- ✅ `github.com/jmoiron/sqlx` (работа с БД)
- ✅ `github.com/lib/pq` (драйвер PostgreSQL)
- ✅ `github.com/golang-migrate/migrate` (миграции)

### 🧪 **Тестовая среда**
- ✅ Готова БД с тестовыми данными
- ✅ Примеры CURL-запросов для проверки API
