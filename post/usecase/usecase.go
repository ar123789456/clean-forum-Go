package usecase

import (
	"forum/models"
	"forum/post"
)

type PostUseCase struct {
	repository post.PostRepository
}

func NewPostUseCase(repository post.PostRepository) *PostUseCase {
	return &PostUseCase{
		repository: repository,
	}
}

func (puc *PostUseCase) Create(user models.Post) error {
	return puc.repository.Create(user)
}

func (puc *PostUseCase) Get(id int64) (models.Post, error) {
	return puc.repository.Get(id)
}
func (puc *PostUseCase) GetAll() ([]models.Post, error) {
	return puc.repository.GetAll()
}

func (puc *PostUseCase) GetByLike() ([]models.Post, error) {
	return puc.repository.GetByLike()
}
func (puc *PostUseCase) GetUserPost(id int64) ([]models.Post, error) {
	return puc.repository.GetUserPost(id)
}
