# Test work for vk

небольшое веб-приложение для заметок, которое включает в себя возможность регистрации пользователей с использованием JWT-токена.

Проект состоит из трёх сервисов: веб-приложения (на React + TypeScript) и двух сервисов на Go (Gin) для аутентификации и выполнения операций CRUD над заметками.

В качестве системы баз данных используется Postgres, инициализация которой производится с помощью файла init.sql.

Для запуска приложения используйте следующую команду docker-compose:

```
docker-compose up --build
```

# Приложение

Изначально пользователь видит все заметки, но не может их редактировать. После регистрации и входа в систему, пользователь получает JWT-токен, который сохраняется в localStorage. После этого, страницы /login и /register перенаправляют на главную страницу.

После входа в систему, пользователь может просматривать свои заметки во вкладке "Мои заметки", каждую из которых можно редактировать или удалять. Также появляется кнопка для добавления новой заметки и кнопка выхода из системы.

Доступна базовая фильтрация по времени создания заметок: за последний день, 3 дня и месяц.

К сожалению, я не успел добавить валидацию логина и пароля. 

## Frontend

Мои знания в области фронтенд-разработки ограничены, поэтому код может выглядеть не идеально, так как я учился в процессе написания сайта. Это касается и дизайна сайта.

API к сервисам прописывается в config.ts. Из-за проблем с CORS лучше было бы использовать прокси-сервер (load balancer), но в качестве временного решения я добавил заголовок CORS к каждому бэкенд-сервису. Его можно модифицировать в docker-compose и в internal/controllers/router.go.

## Backend 

Каждый сервис имеет конфигурационный файл в config/config.yml. Каждый параметр также можно задать через переменные окружения. Подключение к серверам в Docker осуществляется именно таким образом.

Сервис аутентификации работает на порту 8082, а сервис заметок — на порту 8081 (см. файл конфигурации).

Для документации API используется Swagger. Доступ к документации можно получить, перейдя по адресу url/swagger/index.html.

## Где комментарии к коду?

Я придерживаюсь подхода к написанию самодокументируемого кода и не вижу смысла добавлять комментарии к каждому методу просто ради наличия комментариев. В целом, код простой, понятный и легко читаемый, кроме, возможно, фронтенда.

## Тесты

Их нет, не успел
