package usecase

import (
	"forum/comment"
	"forum/models"
)

type Usecase struct {
	repository comment.CommentRepository
}

func NewUsecase(repository comment.CommentRepository) *Usecase {
	return &Usecase{
		repository: repository,
	}
}

func (uc *Usecase) Create(comment models.Comment) ([]models.Comment, error) {
	err := uc.repository.Create(comment)
	if err != nil {
		return []models.Comment{}, err
	}
	return uc.repository.GetAll(comment.PostId)
}
func (uc *Usecase) GetAll(postid int) ([]models.Comment, error) {
	return uc.repository.GetAll(postid)
}
func (uc *Usecase) Delete(commentId, post_id int) ([]models.Comment, error) {
	err := uc.repository.Delete(commentId)
	if err != nil {
		return []models.Comment{}, err
	}
	return uc.repository.GetAll(post_id)
}
