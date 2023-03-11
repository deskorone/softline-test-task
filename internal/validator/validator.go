package validator

import (
	"errors"
	"fmt"
	"regexp"
	"softline-test-task/internal/entity"
	"strings"
)

// Validator - структура имплементирующая UserValidator
type Validator struct {
}

func NewValidator() *Validator {
	return &Validator{}
}

func (v Validator) Validate(user entity.User) error {
	var builder strings.Builder
	if !v.emailValidate(user.Email) {
		builder.WriteString(fmt.Sprintf("email = %s ", user.Email))
	}
	if !v.loginValidate(user.Login) {
		builder.WriteString(fmt.Sprintf("login = %s ", user.Login))
	}
	if !v.phoneValidate(user.PhoneNumber) {
		builder.WriteString(fmt.Sprintf("phone_number = %s ", user.PhoneNumber))
	}
	if !v.passwordValidate(user.Password) {
		builder.WriteString(fmt.Sprintf("password = %s ", user.Password))
	}
	if builder.Len() != 0 {
		return errors.New(builder.String())
	}
	return nil
}

func (v Validator) phoneValidate(phone string) bool {
	r := `^[\+]?[(]?[0-9]{3}[)]?[-\s\.]?[0-9]{3}[-\s\.]?[0-9]{4,6}$`
	check, _ := regexp.Match(r, []byte(phone))
	return check
}

func (v Validator) emailValidate(email string) bool {
	r := `^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`
	check, _ := regexp.Match(r, []byte(email))
	return check
}

func (v Validator) passwordValidate(password string) bool {
	return len(password) > 4
}

func (v Validator) loginValidate(login string) bool {
	return len(login) > 4
}
