package repo

import (
	"github.com/DATA-DOG/go-sqlmock"
	"softline-test-task/internal/entity"
	"testing"
)

func TestCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "login", "email", "password", "phone_number"}).
		AddRow(1, "2", "3", "4", "5")

	u := entity.User{
		Id:          1,
		Login:       "2",
		Email:       "3",
		Password:    "4",
		PhoneNumber: "5",
	}

	mock.ExpectQuery(`insert into users`).WithArgs(u.Login, u.Email, u.Password, u.PhoneNumber).WillReturnRows(rows)

	repo := NewRepository(db)

	_, err = repo.CreateUser(u)
	if err != nil {
		t.Error(err)
	}

}

func TestGetUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "login", "email", "password", "phone_number"}).AddRow("1", "2", "3", "4", "5")

	mock.ExpectQuery(`select (.+) from users`).WithArgs("2").WillReturnRows(rows)

	repo := NewRepository(db)

	_, err = repo.GetUser("2")

	if err != nil {
		t.Error(err)
	}
}

func TestCheckUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	t.Run("Пользователь существует", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"exists"}).AddRow("t")

		login, email, phone := "test", "test", "test"

		mock.ExpectQuery(`select exists()`).WithArgs(login, email, phone).WillReturnRows(rows)

		repo := NewRepository(db)

		check, err := repo.CheckUser(login, email, phone)
		if err != nil {
			t.Error(err)
		}
		if !check {
			t.Error("пользователь должен быть в таблице")
		}
	})

	t.Run("Пользователь не существует", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"exists"}).AddRow("f")

		login, email, phone := "test", "test", "test"

		mock.ExpectQuery(`select exists()`).WithArgs(login, email, phone).WillReturnRows(rows)

		repo := NewRepository(db)

		check, err := repo.CheckUser(login, email, phone)
		if err != nil {
			t.Error(err)
		}
		if check {
			t.Error("пользователь должен быть в таблице")
		}
	})

}
