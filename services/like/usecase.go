package like

import "forum/domain"

type LikeUseCase interface {
	GetPostLike(id, userid int) (domain.Like, error)
	GetCommentLike(id, userid int) (domain.Like, error)
	LikePost(idPost, idUser int, like bool) (domain.Like, error)
	LikeComment(idPost, idUser int, like bool) (domain.Like, error)
}
