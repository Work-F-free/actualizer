# Сервис Актуализатор

## Предназначение
Сервис предназначен для актуализации информации о бронированиях пользователей в системе

## Механизм работы

В сервисе реализовано 2 API-метода, посредством которых производится управление сервисом:

#### Первый метод:
```http request
    GET [your_host]/actualizer/service/start
```
Запускает актуализатор на стороне сервера. Актуализация данных происходит раз в 5 минут. 

В случае, когда время бронирования истекает, актуализатор изменяет статус бронирования на expired (значение поля `expired` становится равно `true`)

#### Второй метод:
```http request
    GET [your_host]/actualizer/service/stop
```
Останавливает работу актуализатора

## Что еще?

В кодовой базе реализована работа репозитория с сущностью `Booking`. 

## Полный список CRUD-операций:

### GetAll
##### Описание
_Получение всех записей о бронировании._
**_Возможна фильтрация по параметрам_**
##### Синтаксис
```go
    GetAll(ctx, params)
```
##### Параметры на вход
- **сtx** - _context.Context_ - Http Context
- **params** - _dto.RequestParams_ - Параметры Query
##### Результат выполнения
- **_[ ]dto.Booking_**
- **_error_**

### GetById
##### Описание
Получение информации о бронирование по его идентификатору.
##### Синтаксис
```go
    GetById(ctx, uuid)
```
##### Параметры на вход
##### Параметры на вход
- **ctx** - _context.Context_ - Http Context
- **uuid** - _uuid.UUID_ - Идентификатор бронирования
##### Результат выполнения
- **_dto.Booking_**
- **_error_**
### Delete
##### Описание
Удаление бронирования.
##### Синтаксис
```go
    Delete(ctx, uuid)
```
##### Параметры на вход
- **ctx** - _context.Context_ - Http Context
- **uuid** - _uuid.UUID_ - Идентификатор бронирования
##### Результат выполнения
- **_error_**


### Update
##### Описание
Обновление бронирования.
##### Синтаксис
```go
    Update(ctx, Booking, uuid)
```
##### Параметры на вход
- **ctx** - _context.Context_ - Http Context
- **Booking** - _*dto.Booking_ - Сущность бронирования
- **uuid** - _uuid.UUID_ - Идентификатор бронирования
##### Результат выполнения
- **_error_**

### Create
##### Описание
Создание бронирования.
##### Синтаксис
```go
    Create(ctx, Booking)
```
##### Параметры на вход
- **сtx** - _context.Context_ - Http Context
- **Booking** - _*dto.Booking_ - Сущность бронирования
##### Результат выполнения
- **_error_** 

# Запуск окружения

1. Создать файл .env по примеру [.env.example](.env.example)
2. Запустить команду:

### Docker:
```bash
    docker-compose up -d
```

### Makefile:
```bash
    make create-app
```

3. Для перезапуска сервиса:

### Docker:
```bash
    docker-compose up -d --build
```

### Makefile:
```bash
    make restart-app
```