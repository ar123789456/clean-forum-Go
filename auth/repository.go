package auth

import "forum/models"

type UserRepository interface {
	Create(models.User) error
	Get(string, string) (*models.User, error)
	Update(models.User) error
	GetByToken(string) (*models.User, error)
}
