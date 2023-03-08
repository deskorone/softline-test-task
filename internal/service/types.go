package service

import "errors"

var (
	ErrUserAlreadyRegistered = errors.New("пользователь с такими данными уже зарегистрирован")
	ErrIncorrectPassword     = errors.New("введенный пароль неверный")
	ErrUserNotFound          = errors.New("пользователь с введенными данными не существует")
	ErrValidation            = errors.New("ошибка валидации")
	ErrHash                  = errors.New("ошибка при хэшировании пароля")
	ErrDataBaseWriting       = errors.New("ошибка записи в базу данных")
)
