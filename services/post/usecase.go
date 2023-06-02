package post

import "forum/domain"

type PostUseCase interface {
	Create(domain.Post) error
	Get(id int64) (domain.Post, error)
	GetPosts(limit int, offset int, cat string) ([]domain.Post, error)
	GetByLike(userid int) ([]domain.Post, error)
	GetUserPost(id int64) ([]domain.Post, error)
}
