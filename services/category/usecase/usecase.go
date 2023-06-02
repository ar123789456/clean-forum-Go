package usecase

import (
	"forum/domain"
	"forum/services/category"
)

type Usecase struct {
	repository category.Repository
}

func NewUsecase(repository category.Repository) *Usecase {
	return &Usecase{
		repository: repository,
	}
}

func (uc *Usecase) Create(name, description string) error {
	return uc.repository.Create(name, description)
}
func (uc *Usecase) Get() ([]domain.Category, error) {
	return uc.repository.Get()
}
