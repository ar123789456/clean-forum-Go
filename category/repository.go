package category

import "forum/models"

type Repository interface {
	Create(name, description string) error
	Get() ([]models.Category, error)
}
