package token

import (
	"github.com/golang-jwt/jwt"
	"softline-test-task/internal/entity"
	"time"
)

/*
	Данный пакет нужен просто для примера результата авторизации
*/

type JwtToken struct {
	secretWord  string
	timeExpired time.Duration
}

func NewAuthToken(secretWord string, expired time.Duration) *JwtToken {
	return &JwtToken{
		secretWord:  secretWord,
		timeExpired: expired,
	}
}

type customClaims struct {
	jwt.StandardClaims
	UserId int64 `json:"user_id"`
}

func (j *JwtToken) GenerateToken(user entity.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &customClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(j.timeExpired).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: user.Id,
	})

	return token.SignedString([]byte(j.secretWord))
}
