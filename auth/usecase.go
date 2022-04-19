package auth

import (
	"errors"
	"forum/models"
)

var WrongPassword = errors.New("wrong password")

type UserUsecase interface {
	SignIn(username, email, password string) (string, error)
	SignUp(name, email, password string) error
	ParseByToken(string) (*models.User, error)
}
