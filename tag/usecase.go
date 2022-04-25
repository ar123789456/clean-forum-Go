package tag

import (
	"forum/models"
)

type Usecase interface {
	Create(name string) error
	Get() ([]models.Tag, error)
}
