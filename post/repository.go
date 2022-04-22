package post

import "forum/models"

type PostRepository interface {
	Create(models.Post) error
	Get(id int64) (models.Post, error)
	GetAll() ([]models.Post, error)
	GetByLike() ([]models.Post, error)
	GetUserPost(id int64) ([]models.Post, error)
}
