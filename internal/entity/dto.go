package entity

// UserLoginDto - объект который приходит при авторизации пользователя
type UserLoginDto struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// UserRegistrationRespDto - объект который возвращается при регистрации
type UserRegistrationRespDto struct {
	Id          int64  `json:"id"`
	Login       string `json:"login"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

// TokenDto - ответ на авторизацию
type TokenDto struct {
	Token string `json:"token"`
}
