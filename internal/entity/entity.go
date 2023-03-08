package entity

// User - структура сущности пользователя
type User struct {
	Id          int64  `json:"id,omitempty"`
	Login       string `json:"login"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}
