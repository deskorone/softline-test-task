package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"softline-test-task/internal/config"
	"softline-test-task/internal/entity"
	"softline-test-task/pkg/helpers"
)

/*
	Транспортный слой приложения
*/

// IService - интерфейс бизнес-логики приложения
type IService interface {
	Registration(user entity.User) (*entity.UserRegistrationRespDto, error)
	Login(user entity.UserLoginDto) (string, error)
}

type Controller struct {
	service IService
}

func NewController(service IService) *Controller {
	return &Controller{service: service}
}

// Register - POST метод осуществляющий регистрацию
func (c *Controller) Register(w http.ResponseWriter, r *http.Request) {
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
func (c Controller) Login(w http.ResponseWriter, r *http.Request) {
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

// CreateServer - функция для создания http-сервера
func CreateServer(c *Controller, conf config.Server) *http.Server {

	mux := http.NewServeMux()

	mux.HandleFunc("/registration", c.Register)
	mux.HandleFunc("/login", c.Login)

	return &http.Server{
		Addr:    fmt.Sprintf(":%d", conf.Port),
		Handler: mux,
	}
}
