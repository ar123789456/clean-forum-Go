package auth

import (
	"errors"
	"forum/domain"
)

var WrongPassword = errors.New("wrong password")

type UserUsecase interface {
	SignIn(username, email, password string) (string, int, error)
	SignUp(user domain.User) error
	GetByToken(string) (*domain.User, error)
}
