# Marketplace REST API

[![Go](https://img.shields.io/badge/Go-1.24.4-blue.svg)](https://golang.org/)
[![Docker](https://img.shields.io/badge/Docker-Container-blue.svg)](https://www.docker.com/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-17-blue.svg)](https://www.postgresql.org/)

---

## Описание проекта

Marketplace — это REST API приложение условного маркетплейса, реализованное на языке Go с использованием фреймворка Echo. Приложение позволяет пользователям регистрироваться, авторизовываться, размещать объявления и просматривать ленту объявлений с фильтрами и сортировками.

Все данные передаются и принимаются в формате JSON. Проект включает в себя авторизацию на основе JWT, управление пользователями и объявлениями, а также постраничный вывод объявлений.

---

---
## Запуск проекта

Возможны два варианта работы с приложением:

- Рабочее приложение: [https://marketplace-fxdq.onrender.com](https://marketplace-fxdq.onrender.com)

> ⚠️ Из-за ограничений бесплатного тарифа хостинга первое обращение к приложению после длительного простоя (примерно 15 минут) может сопровождаться небольшой задержкой при запуске.

- Запуск локально с помощью Docker и docker-compose

### Требования

- Установленные Docker и Docker Compose.

### Команды запуска

- Запуск через Docker Compose:

```bash
docker-compose -f docker-compose.yml up
```

- Запуск через Makefile (дефолтная команда `make docker`):

```
make
```

Или явно:

```
make docker
```

> При запуске поднимаются контейнеры с PostgreSQL и приложением. Подключение к базе данных, миграции и настройка конфигурации выполняются автоматически.


## Основные возможности

- **Регистрация пользователей** с валидацией логина и пароля
- **Авторизация пользователей** с выдачей JWT токена
- **Размещение новых объявлений** только для авторизованных пользователей
- **Просмотр ленты объявлений** с пагинацией, сортировкой (по дате и цене), фильтрацией по цене
- Возвращение признака принадлежности объявления текущему пользователю (если он авторизован)
- Полное покрытие логированием и обработкой ошибок

---

## Технологии и инструменты

- Язык: **Go**
- Веб-фреймворк: [Echo](https://echo.labstack.com/)
- База данных: **PostgreSQL**
- Миграции: [Goose](https://github.com/pressly/goose)
- Логирование: [Logrus](https://github.com/sirupsen/logrus)
- Аутентификация: **JWT**
- Валидация данных: [go-playground/validator](https://github.com/go-playground/validator)
- Шифрование паролей: [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)
- PostgreSQL драйвер: [pgx](https://github.com/jackc/pgx)
- Контейнеризация: **Docker**, **docker-compose**
- **Makefile** для упрощения запуска и сборки

---

## Структура проекта

```plaintext
├───cmd/marketplace
│   └─── main.go             # Точка входа приложения
├─── configs/                # Конфигурационные файлы
└─── internal/
    ├─── app/                # Создание и запуск приложения
    ├─── auth/               # Вся логика работы с jwt
    ├─── config/             # Инициализация конфига
    ├─── db/                 # Подключение к базе данных
    │   └─── migrations/     # Скрипты миграций
    ├─── handlers/           # HTTP-обработчики (эндпоинты)
    ├─── logger/             # Создание и настройка логера
    ├─── middlewares/        # Middleware для авторизации
    ├─── models/             # Структуры данных
    ├─── services/           # Сервисный слой
    └─── storage/            # Работа с базой данных
├── Dockerfile               # Описание контейнера приложения
├── docker-compose.yml       # Запуск приложения и БД через docker-compose
├── Makefile                 # Команды для сборки и запуска psql 
└── README.md                # Документация проекта
```



## Использование API

### Регистрация пользователя

- **POST** `/register`

```json
{
    "login": "user123@example.com",
    "password": "securePassword1!"
}
```
> Логин обязатльно должен быть в формате электронной почты (с @)

Ответ — данные зарегистрированного пользователя (без пароля).
```json
{
    "id": 2,
    "login": "user123@example.com",
    "created_at": "2025-07-21T23:34:35.503909Z"
}
```


### Авторизация пользователя

- **GET** `/login`

```json
{
    "login": "user123@example.com",
    "password": "securePassword1!"
}
```

Ответ — JWT токен. 
```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTMxNTkwNjMsInN1YiI6Mn0.T7Xv_JOpzExeeua9sb4SNBFzBKt86C1vTgqFDlZoDDU"
}
```

Передавайте его в заголовке запросов:

```
Authorization: Bearer <token>
```


### Размещение объявления

- **POST** `/postAd`

Требуется авторизация.

```json
{
"title": "Заголовок объявления",
"text": "Описание объявления",
"image_url": "https://example.com/image.jpg",
"price": 1000
}
```
Ответ - данные о созданном объявлении

```json
{
    "id": 4,
    "title": "Vans Pro",
    "text": "Отличные кроссовки. Ходить в них очень классно. Удобно невероятно",
    "image_url": "https://soulcycle.com/cdn/shop/files/1cc17d3f0591118bb3b9c7546269b575.jpg",
    "price": 7800.5,
    "author_id": 2,
    "is_mine": true,
    "created_at": "2025-07-21T23:43:54.680969Z"
}
```

### Просмотр ленты объявлений

- **GET** `/getList`


Пример: https://marketplace-fxdq.onrender.com/getList?page=1&limit=10&order=date&direction=asc&min=1000&max=100000


Параметры запроса (query parameters):

- `page` — номер страницы (по умолчанию 1)
- `limit` — количество объявлений на странице (по умолчанию 10)
- `order` — поле сортировки (`date` или `price`)
- `direction` — направление сортировки (`asc` или `desc`)
- `min` — минимальная цена (фильтр)
- `max` — максимальная цена (фильтр)

Ответ — список объявлений с необходимыми полями и признаком принадлежности текущему пользователю (если авторизован).

```json
[
    {
        "id": 1,
        "title": "Пылесос",
        "text": "Отличный пылесос. Убираться с ним очень классно. Сосет невероятно",
        "image_url": "https://bosch-centre.ru/upload/resize_cache/iblock/355/553_440_1/355e1f9ff4feb52868b56295160e5833.jpg",
        "price": 10000,
        "author_id": 1,
        "is_mine": true,
        "created_at": "0001-01-01T00:00:00Z"
    },
    {
        "id": 2,
        "title": "Фен Dyson",
        "text": "Отличный фен. Сушить волосы с ним очень классно. Дует невероятно",
        "image_url": "https://bosch-centre.ru/upload/resize_cache/iblock/355/553_440_1/355e1f9ff4feb52868b56295160e5833.jpg",
        "price": 2500,
        "author_id": 1,
        "is_mine": true,
        "created_at": "0001-01-01T00:00:00Z"
    },
    {
        "id": 3,
        "title": "Iphone 16 Pro Max",
        "text": "Отличный телефон. Работать и звонить с ним очень классно. Работает невероятно",
        "image_url": "https://www.apple.com/newsroom/images/2024/09/apple-debuts-iphone-16-pro-and-iphone-16-pro-max/article/Apple-iPhone-16-Pro-hero-240909_inline.jpg.large.jpg",
        "price": 90000,
        "author_id": 1,
        "is_mine": true,
        "created_at": "0001-01-01T00:00:00Z"
    },
    {
        "id": 4,
        "title": "Vans Pro",
        "text": "Отличные кроссовки. Ходить в них очень классно. Удобно невероятно",
        "image_url": "https://soulcycle.com/cdn/shop/files/1cc17d3f0591118bb3b9c7546269b575.jpg",
        "price": 7800.5,
        "author_id": 2,
        "is_mine": false,
        "created_at": "0001-01-01T00:00:00Z"
    }
]
```


## Контакты и ссылки
Если у вас есть вопросы или предложения, пожалуйста, свяжитесь со мной.
- Электронная почта: [zhirenkoartem@gmail.com](mailto:zhirenkoartem@gmail.com)

- Telegram : [@teema_0](https://t.me/teema_0)
