package comment

import "forum/domain"

type CommentUseCase interface {
	Create(domain.Comment) ([]domain.Comment, error)
	GetAll(postid int) ([]domain.Comment, error)
	Delete(commentId, postid int) ([]domain.Comment, error)
}
