package usecase

import (
	"forum/domain"
	"forum/services/comment"
)

type Usecase struct {
	repository comment.CommentRepository
}

func NewUsecase(repository comment.CommentRepository) *Usecase {
	return &Usecase{
		repository: repository,
	}
}

func (uc *Usecase) Create(comment domain.Comment) ([]domain.Comment, error) {
	err := uc.repository.Create(comment)
	if err != nil {
		return []domain.Comment{}, err
	}
	return uc.repository.GetAll(comment.PostId)
}
func (uc *Usecase) GetAll(postid int) ([]domain.Comment, error) {
	return uc.repository.GetAll(postid)
}
func (uc *Usecase) Delete(commentId, post_id int) ([]domain.Comment, error) {
	err := uc.repository.Delete(commentId)
	if err != nil {
		return []domain.Comment{}, err
	}
	return uc.repository.GetAll(post_id)
}
