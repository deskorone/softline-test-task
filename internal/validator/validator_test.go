package validator

import (
	"softline-test-task/internal/entity"
	"testing"
)

func TestValidate(t *testing.T) {

	t.Run("Валидные данные", func(t *testing.T) {
		newValidator := NewValidator()

		user := entity.User{
			Id:          0,
			Login:       "exampleLogin",
			Email:       "email@example.ru",
			Password:    "123456",
			PhoneNumber: "88005555555",
		}

		if err := newValidator.Validate(user); err != nil {
			t.Error(err)
		}
	})

	t.Run("Данные невалидны", func(t *testing.T) {
		newValidator := NewValidator()

		cases := []entity.User{
			{
				Id:          1,
				Login:       "1",
				Email:       "example@mail.ru",
				Password:    "1234",
				PhoneNumber: "88005555555",
			}, {
				Id:          2,
				Login:       "logintest",
				Email:       "exaWWW",
				Password:    "123789",
				PhoneNumber: "88005555555",
			}, {
				Id:          3,
				Login:       "qtesttest",
				Email:       "example@mail.ru",
				Password:    "a",
				PhoneNumber: "88005555555",
			}, {
				Id:          3,
				Login:       "qtesttest",
				Email:       "example@mail.ru",
				Password:    "password123",
				PhoneNumber: "8",
			},
		}

		for _, user := range cases {
			if err := newValidator.Validate(user); err == nil {
				t.Error("Должна была вернуться ошибка валидации")
			}
		}
	})
}
