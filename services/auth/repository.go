package auth

import "forum/domain"

type UserRepository interface {
	Create(domain.User) error
	Get(string, string) (*domain.User, error)
	Update(domain.User) error
	GetByToken(string) (*domain.User, error)
}
