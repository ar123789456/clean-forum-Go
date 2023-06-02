package category

import "forum/domain"

type Repository interface {
	Create(name, description string) error
	Get() ([]domain.Category, error)
}
