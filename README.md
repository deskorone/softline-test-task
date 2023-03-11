# Краткое описание по api

## Запустить docker-compose можно командой ```make up```

### Сервер запускается на порту 8080

## /register - Post запрос на регистрацию пользователя

### Тело запроса
```json 
{
  "login" : "testtest",
  "email" : "test@example.com",
  "password" : "secret_word",
  "phone_number" : "89005552525"
}
```
### Ответ на запрос
```json
{
  "id" : 1,
  "login" : "testtest",
  "email" : "test@example.com",
  "phone_number" : "89005552525"
}
```

## /login - Post запрос на авторизацию пользователя
### Тело запроса
```json
{
  "login" : "testtest",
  "password" : "secret_word"
}
```
### Ответ на запрос при успешной авторизации
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6..."
}
```
