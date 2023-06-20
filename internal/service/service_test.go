package service

import (
	"errors"
	"github.com/golang/mock/gomock"
	"softline-test-task/internal/entity"
	mock_service "softline-test-task/mocks"
	"testing"
)

func TestRegistration(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	authToken := mock_service.NewMockAuthToken(ctrl)
	repo := mock_service.NewMockUserRepository(ctrl)
	validator := mock_service.NewMockUserValidator(ctrl)
	hasher := mock_service.NewMockHasher(ctrl)
	service := NewService(repo, hasher, authToken, validator)

	t.Run("Успешная регистрация", func(t *testing.T) {
		user := entity.User{
			Id:          0,
			Login:       "1",
			Email:       "123@example.ru",
			Password:    "12345",
			PhoneNumber: "88000003333",
		}

		validator.EXPECT().Validate(user).Return(nil).Times(1)
		repo.EXPECT().CheckUser(user.Login, user.Email, user.PhoneNumber).Return(false, nil).Times(1)
		hasher.EXPECT().Hash(user.Password).Return(user.Password, nil).Times(1)
		repo.EXPECT().CreateUser(user).Return(nil, nil)

		if _, err := service.Registration(user); err != nil {
			t.Error(err)
		}
	})

	t.Run("Невалидные данные", func(t *testing.T) {
		cases := []entity.User{
			{
				Id:          1,
				Login:       "1",
				Email:       "example@mail.ru",
				Password:    "",
				PhoneNumber: "",
			}, {
				Id:          2,
				Login:       "asdqw",
				Email:       "exaWWW",
				Password:    "123",
				PhoneNumber: "",
			}, {
				Id:          2,
				Login:       "qwe",
				Email:       "example@mail.ru",
				Password:    "asd",
				PhoneNumber: "asdfaf",
			},
		}

		for _, v := range cases {
			validator.EXPECT().Validate(v).Return(ErrValidation).Times(1)
			_, err := service.Registration(v)
			if !errors.Is(err, ErrValidation) {
				t.Error(err)
			}
		}

	})

	t.Run("Ошибка при хэшировании", func(t *testing.T) {
		user := entity.User{
			Id:          0,
			Login:       "1",
			Email:       "123@example.ru",
			Password:    "12345",
			PhoneNumber: "88000003333",
		}
		validator.EXPECT().Validate(user).Return(nil).Times(1)
		repo.EXPECT().CheckUser(user.Login, user.Email, user.PhoneNumber).Return(false, nil).Times(1)
		hasher.EXPECT().Hash(user.Password).Return(user.Password, errors.New("ошибка хэширования")).Times(1)

		if _, err := service.Registration(user); !errors.Is(err, ErrHash) {
			t.Error(err)
		}
	})

	t.Run("Ошибка записи в бд", func(t *testing.T) {
		user := entity.User{
			Id:          0,
			Login:       "1",
			Email:       "123@example.ru",
			Password:    "12345",
			PhoneNumber: "88000003333",
		}
		validator.EXPECT().Validate(user).Return(nil).Times(1)
		repo.EXPECT().CheckUser(user.Login, user.Email, user.PhoneNumber).Return(false, nil).Times(1)
		hasher.EXPECT().Hash(user.Password).Return(user.Password, nil).Times(1)
		repo.EXPECT().CreateUser(user).Return(nil, errors.New("")).Times(1)

		if _, err := service.Registration(user); !errors.Is(err, ErrDataBaseWriting) {
			t.Error(err)
		}
	})

	t.Run("Пользователь уже зарегистрирован", func(t *testing.T) {
		user := entity.User{
			Id:          0,
			Login:       "1",
			Email:       "123@example.ru",
			Password:    "12345",
			PhoneNumber: "88000003333",
		}

		validator.EXPECT().Validate(user).Return(nil).Times(1)
		repo.EXPECT().CheckUser(user.Login, user.Email, user.PhoneNumber).Return(true, nil).Times(1)

		if _, err := service.Registration(user); !errors.Is(err, ErrUserAlreadyRegistered) {
			t.Error(err)
		}

	})
}

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	authToken := mock_service.NewMockAuthToken(ctrl)
	repo := mock_service.NewMockUserRepository(ctrl)
	validator := mock_service.NewMockUserValidator(ctrl)
	hasher := mock_service.NewMockHasher(ctrl)
	service := NewService(repo, hasher, authToken, validator)

	t.Run("Пользователь не найден", func(t *testing.T) {
		req := entity.UserLoginDto{
			Login:    "12333",
			Password: "12333",
		}

		repo.EXPECT().GetUser(req.Login).Return(nil, errors.New(""))

		if _, err := service.Login(req); !errors.Is(err, ErrUserNotFound) {
			t.Error(err)
		}

		repo.EXPECT().GetUser(req.Login).Return(nil, nil)

		if _, err := service.Login(req); !errors.Is(err, ErrUserNotFound) {
			t.Error(err)
		}
	})

	t.Run("Пароль неверный", func(t *testing.T) {
		req := entity.UserLoginDto{
			Login:    "12333",
			Password: "12333",
		}

		res := entity.User{
			Id:          1,
			Login:       "12333",
			Email:       "12333",
			Password:    "12333",
			PhoneNumber: "12333",
		}

		repo.EXPECT().GetUser(req.Login).Return(&res, nil)
		hasher.EXPECT().CheckHash(req.Password, res.Password).Return(false)

		if _, err := service.Login(req); !errors.Is(err, ErrIncorrectPassword) {
			t.Error(err)
		}
	})
}
