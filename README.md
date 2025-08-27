# Task Manager

## Контейнеризация

Для упрощения развертывания приложения используется Docker. 

### 1. Сборка Docker-образа
```bash
docker build -t task-manager .
```

### 2. Запуск контейнера
```bash
docker run -p 8080:8080 --env-file .env task-manager
```

### 3. Использование Docker Compose
Для запуска приложения и базы данных одновременно используйте `docker-compose.yml`:
```bash
docker-compose up
```

---

## Установка и запуск вручную

### 1. Установите зависимости
Убедитесь, что на вашем компьютере установлены:
- Go (версия 1.20 или выше)
- PostgreSQL
- Docker (опционально, для контейнеризации)

### 2. Настройте базу данных
Создайте базу данных `task_manager` в PostgreSQL:
```sql
CREATE DATABASE task_manager;
```

Примените миграции из папки [`migrations`](migrations ) для создания таблиц:
```bash
go run cmd/app/main.go
```

### 3. Настройте переменные окружения
Создайте файл [`.env`](.env ) в корне проекта и укажите параметры подключения к базе данных:
```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=task_manager
SSL_MODE=disable
```

### 4. Запустите приложение
Запустите приложение с помощью команды:
```bash
go run cmd/app/main.go
```

Приложение будет доступно по адресу: `http://localhost:8080`.

---

## API-эндпоинты

### 1. Регистрация пользователя
**POST** `/register`  
**Тело запроса:**
```json
{
  "username": "example_user",
  "password": "example_password"
}
```
**Ответ:**
```json
{
  "message": "User registered successfully"
}
```

### 2. Логин пользователя
**POST** `/login`  
**Тело запроса:**
```json
{
  "username": "example_user",
  "password": "example_password"
}
```
**Ответ:**
```json
{
  "message": "Login successful"
}
```

### 3. Работа с задачами
- **Создание задачи:**
  **POST** `/tasks`  
  **Тело запроса:**
  ```json
  {
    "user_id": 1,
    "title": "Complete report",
    "description": "Write the final report for the internship"
  }
  ```
  **Ответ:**
  ```json
  {
    "message": "Task created successfully",
    "task": {
      "id": 1,
      "title": "Complete report",
      "description": "Write the final report for the internship",
      "status": "pending",
      "created_at": "2025-08-23T12:00:00Z"
    }
  }
  ```

- **Получение задач:**
  **GET** `/tasks`  
  **Ответ:**
  ```json
  [
    {
      "id": 1,
      "user_id": 1,
      "title": "Complete report",
      "description": "Write the final report for the internship",
      "status": "pending",
      "created_at": "2025-08-23T12:00:00Z"
    }
  ]
  ```

- **Обновление задачи:**
  **PUT** `/tasks/{id}`  
  **Тело запроса:**
  ```json
  {
    "title": "Updated title",
    "description": "Updated description",
    "status": "completed"
  }
  ```

- **Удаление задачи:**
  **DELETE** `/tasks/{id}`  
  **Ответ:**
  ```json
  {
    "message": "Task deleted successfully"
  }
  ```

---

