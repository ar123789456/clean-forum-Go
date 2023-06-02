package like

import "forum/domain"

type LikeRepository interface {
	GetPostLike(id, userid int) (domain.Like, error)
	GetCommentLike(id, userid int) (domain.Like, error)
	LikePostUpdate(idPost, idUser, like int) error
	LikeCommentUpdate(idPost, idUser, like int) error
	LikePost(idPost, idUser, like int) error
	LikeComment(idPost, idUser, like int) error
}
