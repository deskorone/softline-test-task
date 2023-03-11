package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"softline-test-task/internal/entity"
	"softline-test-task/pkg/helpers"
)

/*
	Транспортный слой приложения
*/

// AuthService - интерфейс бизнес-логики приложения
type AuthService interface {
	Registration(user entity.User) (*entity.UserRegistrationRespDto, error)
	Login(user entity.UserLoginDto) (string, error)
}

type RestController struct {
	service AuthService
}

func NewRestController(service AuthService) *RestController {
	return &RestController{service: service}
}

// Register - POST метод осуществляющий регистрацию
func (c *RestController) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helpers.Failed(w, helpers.MethodNotAllowed, http.StatusMethodNotAllowed)
		return
	}
	user := entity.User{}

	// Парсим тело запроса
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		helpers.Failed(w, fmt.Sprintf("Ошибка при конвертировании: %v", err), http.StatusBadRequest)
		return
	}

	// Передаем данные на уровень бизнес-логики
	resp, err := c.service.Registration(user)
	if err != nil {
		helpers.Failed(w, fmt.Sprintf("Произоошла ошибка при регистрации: %v", err), http.StatusBadRequest)
		return
	}

	helpers.Success(w, resp)
}

// Login - POST метод осуществляющий авторизацию
func (c RestController) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helpers.Failed(w, helpers.MethodNotAllowed, http.StatusMethodNotAllowed)
		return
	}
	user := entity.UserLoginDto{}

	// Парсим тело запроса
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		helpers.Failed(w, fmt.Sprintf("Ошибка при конвертировании: %v", err), http.StatusBadRequest)
		return
	}

	// Передаем данные на уровень бизнес-логики
	token, err := c.service.Login(user)
	if err != nil {
		helpers.Failed(w, fmt.Sprintf("Ошибка при авторизации: %v", err), http.StatusBadRequest)
		return
	}

	helpers.Success(w, entity.TokenDto{Token: token})
}
