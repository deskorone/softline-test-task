# Краткое описание по api

## Запустить docker-compose можно командой ```make up```

## /register - Post запрос на регистрацию пользователя

### Тело запроса
```json 
{
  "login" : "test",
  "email" : "test@example.com",
  "password" : "secret word",
  "phone_number" : "780000000"
}
```
### Ответ на запрос
```json
{
  "id" : 1,
  "login" : "test",
  "email" : "test@example.com",
  "phone_number" : "780000000"
}
```

## /login - Post запрос на авторизацию пользователя
### Тело запроса
```json
{
  "login" : "test",
  "password" : "secret_word"
}
```
### Ответ на запрос
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6..."
}
```
