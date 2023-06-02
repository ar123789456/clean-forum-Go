package usecase

import (
	"forum/domain"
	"forum/services/tag"
)

type Usecase struct {
	repository tag.Repository
}

func NewUsecase(repository tag.Repository) *Usecase {
	return &Usecase{
		repository: repository,
	}
}

func (uc *Usecase) Create(name string) error {
	return uc.repository.Create(name)
}
func (uc *Usecase) Get() ([]domain.Tag, error) {
	return uc.repository.Get()
}
