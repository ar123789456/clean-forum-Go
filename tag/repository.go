package tag

import "forum/models"

type Repository interface {
	Create(name string) error
	Get() ([]models.Tag, error)
}
