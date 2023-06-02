package category

import (
	"forum/domain"
)

type Usecase interface {
	Create(name, description string) error
	Get() ([]domain.Category, error)
}
