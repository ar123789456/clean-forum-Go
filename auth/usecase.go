package auth

import "forum/models"

type UserUsecase interface {
	SignIn(string, string) error
	SignUp(string, string) (string, error)
	ParseToken(string) (*models.User, error)
}
