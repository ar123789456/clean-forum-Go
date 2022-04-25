package category

import (
	"forum/models"
)

type Usecase interface {
	Create(name, description string) error
	Get() ([]models.Category, error)
}
