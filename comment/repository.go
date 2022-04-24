package comment

import "forum/models"

type CommentRepository interface {
	Create(models.Comment) error
	GetAll(postid int) ([]models.Comment, error)
	Delete(commentId int) error
}
