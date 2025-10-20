# AI Chat Go

Go реализация AiModelController из PHP проекта ai-chat.

## Структура проекта

```
.
├── config/          # Конфигурация приложения
├── database/        # Подключение к БД и миграции
├── handlers/        # HTTP обработчики (контроллеры)
├── models/          # Модели данных
├── main.go          # Точка входа
├── go.mod           # Зависимости Go
└── .env.example     # Пример конфигурации
```

## Установка

1. Клонируйте репозиторий
2. Скопируйте `.env.example` в `.env` и настройте параметры БД:
   ```bash
   cp .env.example .env
   ```
3. Установите зависимости:
   ```bash
   go mod download
   ```

## Настройка базы данных

Приложение использует PostgreSQL. Создайте базу данных:

## Запуск

```bash
go run main.go
```

Или соберите и запустите:

```bash
go build -o ai-chat-go
./ai-chat-go
```

## API Endpoints

### GET /go-api/models

Возвращает список активных AI моделей.

**Response:**
```json
{
  "data": [
    {
      "id": 1,
      "uuid": "550e8400-e29b-41d4-a716-446655440000",
      "name": "GPT-4",
      "group": "natural_language_processing",
      "task": "text_generation",
      "driver": "localAi",
      "model_name": "gpt-4",
      "active": true
    }
  ]
}
```

## Переменные окружения

- `DB_HOST` - Хост базы данных (по умолчанию: localhost)
- `DB_PORT` - Порт базы данных (по умолчанию: 5432)
- `DB_USER` - Пользователь БД (по умолчанию: postgres)
- `DB_PASSWORD` - Пароль БД (по умолчанию: password)
- `DB_NAME` - Имя базы данных (по умолчанию: ai_chat)
- `SERVER_PORT` - Порт сервера (по умолчанию: 8080)

## Модель данных

### AiModel

| Поле | Тип | Описание |
|------|-----|----------|
| id | uint | Первичный ключ |
| uuid | string | UUID модели (генерируется автоматически) |
| name | string | Название модели |
| group | string | Группа модели |
| task | string | Задача модели |
| driver | string | Драйвер модели |
| model_name | string | Имя модели в драйвере |
| host | string | Хост API (опционально) |
| token | string | Токен доступа (опционально) |
| description | string | Описание модели (опционально) |
| active | bool | Активна ли модель |
| created_at | time.Time | Дата создания |
| updated_at | time.Time | Дата обновления |

## Технологии

- **Gin** - HTTP веб-фреймворк
- **GORM** - ORM для работы с БД
- **PostgreSQL** - База данных
- **UUID** - Генерация уникальных идентификаторов
