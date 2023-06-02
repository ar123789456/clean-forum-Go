package post

import (
	"forum/domain"
)

type PostRepository interface {
	Create(domain.Post) error
	Get(id int64) (domain.Post, error)
	GetAll() ([]domain.Post, error)
	GetByLike(userid int) ([]domain.Post, error)
	GetUserPost(id int64) ([]domain.Post, error)
}
