package comment

import "forum/domain"

type CommentRepository interface {
	Create(domain.Comment) error
	GetAll(postid int) ([]domain.Comment, error)
	Delete(commentId int) error
}
