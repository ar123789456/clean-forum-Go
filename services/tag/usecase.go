package tag

import (
	"forum/domain"
)

type Usecase interface {
	Create(name string) error
	Get() ([]domain.Tag, error)
}
