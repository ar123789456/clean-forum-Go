package like

import "forum/models"

type LikeRepository interface {
	GetPostLike(id, userid int) (models.Like, error)
	GetCommentLike(id, userid int) (models.Like, error)
	LikePostUpdate(idPost, idUser, like int) error
	LikeCommentUpdate(idPost, idUser, like int) error
	LikePost(idPost, idUser, like int) error
	LikeComment(idPost, idUser, like int) error
}
