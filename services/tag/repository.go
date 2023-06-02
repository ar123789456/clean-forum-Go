package tag

import "forum/domain"

type Repository interface {
	Create(name string) error
	Get() ([]domain.Tag, error)
}
