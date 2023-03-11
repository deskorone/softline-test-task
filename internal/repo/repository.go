package repo

import (
	"database/sql"
	"fmt"
	"log"
	"softline-test-task/internal/config"
	"softline-test-task/internal/entity"
)

// OpenConnection - функция для открытия соединения с бд
func OpenConnection(conf config.DatabaseConfig) (*sql.DB, error) {
	host, port, user, password, dbname := conf.Host, conf.Port, conf.User, conf.Password, conf.DbName
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return db, nil
}

// NewRepository - функция-конструктор для создания Repository
func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

// Repository - структура для работы с бд имплементирует UserRepository
type Repository struct {
	db *sql.DB
}

func (r *Repository) CheckUser(login, email, phone string) (bool, error) {
	q := ` select exists(select 1 from users where login=$1 or email=$2 or phone_number=$3);`

	var exist bool
	err := r.db.QueryRow(q, login, email, phone).Scan(&exist)
	return exist, err
}

// CreateUser - метод для создания пользователя
func (r *Repository) CreateUser(user entity.User) (*entity.UserRegistrationRespDto, error) {
	q := `insert into users (login, email, password, phone_number) values ($1, $2, $3, $4) returning *;`

	err := r.db.QueryRow(q, user.Login, user.Email, user.Password, user.PhoneNumber).
		Scan(&user.Id, &user.Login, &user.Email, &user.Password, &user.PhoneNumber)

	return &entity.UserRegistrationRespDto{
		Id:          user.Id,
		Login:       user.Login,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
	}, err
}

// GetUser - метод для взятия пользователя из бд по логину
func (r *Repository) GetUser(login string) (*entity.User, error) {
	u := entity.User{}

	q := `select id, login, email, password, phone_number  from users u where login=$1 limit 1`
	err := r.db.QueryRow(q, login).Scan(&u.Id, &u.Login, &u.Email, &u.Password, &u.PhoneNumber)
	return &u, err
}
