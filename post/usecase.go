package post

import "forum/models"

type PostUseCase interface {
	Create(models.Post) error
	Get(id int64) (models.Post, error)
	GetAll() ([]models.Post, error)
	GetByLike(userid int) ([]models.Post, error)
	GetUserPost(id int64) ([]models.Post, error)
}
