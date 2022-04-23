package like

import "forum/models"

type LikeUseCase interface {
	GetPostLike(id, userid int) (models.Like, error)
	GetCommentLike(id, userid int) (models.Like, error)
	LikePost(idPost, idUser int, like bool) (models.Like, error)
	LikeComment(idPost, idUser int, like bool) (models.Like, error)
}
