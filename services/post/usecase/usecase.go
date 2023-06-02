package usecase

import (
	"forum/domain"
	"forum/services/post"
)

type PostUseCase struct {
	repository post.PostRepository
}

func NewPostUseCase(repository post.PostRepository) *PostUseCase {
	return &PostUseCase{
		repository: repository,
	}
}

func (puc *PostUseCase) Create(user domain.Post) error {
	return puc.repository.Create(user)
}

func (puc *PostUseCase) Get(id int64) (domain.Post, error) {
	return puc.repository.Get(id)
}
func (puc *PostUseCase) GetPosts(limit int, offset int, cat string) ([]domain.Post, error) {
	if cat == "all" {
		cat = ""
	}
	return puc.repository.GetPosts(limit, offset, cat)
}

func (puc *PostUseCase) GetByLike(userid int) ([]domain.Post, error) {
	return puc.repository.GetByLike(userid)
}
func (puc *PostUseCase) GetUserPost(id int64) ([]domain.Post, error) {
	return puc.repository.GetUserPost(id)
}
