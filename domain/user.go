package domain

import (
	"errors"
	"net/mail"
)

var (
	ErrUserNotFound  = errors.New("user not found")
	ErrWrongPassword = errors.New("wrong password")
	ErrorEmptyEmail  = errors.New("email is empty")
)

type Email string

func NewEmail(e string) (Email, error) {
	if e == "" {
		return "", ErrorEmptyEmail
	}
	_, err := mail.ParseAddress(e)
	return Email(e), err
}

type User struct {
	ID       int64
	UUID     string
	Username string
	Email    Email
	Password string
}
