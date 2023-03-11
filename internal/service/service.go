package service

import (
	"fmt"
	"softline-test-task/internal/entity"
)

/*
	service - слой бизнес-логики
*/

// UserRepository - интерфейс для работы с базой данных
type UserRepository interface {
	CreateUser(user entity.User) (*entity.UserRegistrationRespDto, error) // Метод добавления пользователя в бд
	GetUser(login string) (*entity.User, error)                           // Метод получения пользователя из бд
	CheckUser(login, email, phone string) (bool, error)                   // Проверка на запись пользователя в базе
}

// Hasher - интерфейс для хэширования паролей
type Hasher interface {
	Hash(password string) (string, error) // Метод хэширующий пароль
	CheckHash(hash, password string) bool // Метод сопоставляющий хэш и пароль
}

// AuthToken - интерфейс для работы с токенами авторизации (для примера результата авторизации)
type AuthToken interface {
	GenerateToken(user entity.User) (string, error) // Метод для генерации токена авторизации
}

// UserValidator - интерфейс для валидации пользователя
type UserValidator interface {
	Validate(user entity.User) error
}

// Service - структура реализующая бизнес-логику приложения
type Service struct {
	storage   UserRepository
	hasher    Hasher
	authToken AuthToken
	validator UserValidator
}

// NewService - функция для получения экземпляра сервиса
func NewService(s UserRepository, h Hasher, a AuthToken, v UserValidator) *Service {
	return &Service{storage: s, hasher: h, authToken: a, validator: v}
}

// Registration - бизнес-логика регистрации
func (s *Service) Registration(user entity.User) (*entity.UserRegistrationRespDto, error) {
	if err := s.validator.Validate(user); err != nil {
		return nil, fmt.Errorf("%w - %v", ErrValidation, err)
	}

	check, err := s.storage.CheckUser(user.Login, user.Email, user.PhoneNumber)
	if err != nil {
		return nil, err
	}
	if check {
		return nil, ErrUserAlreadyRegistered
	}
	// Если все успешно, то хэшируем пароль и присваиваем результат пришедшей структуре
	hashedPassword, err := s.hasher.Hash(user.Password)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrHash, err)
	}
	user.Password = hashedPassword

	// Добавляем пользователя в базу данных
	res, err := s.storage.CreateUser(user)
	if err != nil {
		return nil, fmt.Errorf("%w : %v", ErrDataBaseWriting, err)
	}
	return res, nil
}

// Login - бизнес-логика авторизации
func (s *Service) Login(user entity.UserLoginDto) (string, error) {
	u, err := s.storage.GetUser(user.Login)
	if err != nil {
		return "", fmt.Errorf("%w : %v", ErrUserNotFound, err)
	}
	if u == nil {
		return "", ErrUserNotFound
	}

	check := s.hasher.CheckHash(u.Password, user.Password)
	if !check {
		return "", ErrIncorrectPassword
	}

	return s.authToken.GenerateToken(*u)
}
