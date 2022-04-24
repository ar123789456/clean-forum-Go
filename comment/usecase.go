package comment

import "forum/models"

type CommentUseCase interface {
	Create(models.Comment) ([]models.Comment, error)
	GetAll(postid int) ([]models.Comment, error)
	Delete(commentId, postid int) ([]models.Comment, error)
}
