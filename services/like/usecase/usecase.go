package usecase

import (
	"forum/domain"
	"forum/services/like"
	"log"
)

type LikeUseCase struct {
	repository like.LikeRepository
}

func NewLikeUseCase(repository like.LikeRepository) *LikeUseCase {
	return &LikeUseCase{
		repository: repository,
	}
}

func (luc *LikeUseCase) GetPostLike(id, userid int) (domain.Like, error) {
	return luc.repository.GetPostLike(id, userid)
}

func (luc *LikeUseCase) GetCommentLike(id, userid int) (domain.Like, error) {
	return luc.repository.GetCommentLike(id, userid)
}

func (luc *LikeUseCase) LikePost(idPost, idUser int, like bool) (domain.Like, error) {
	likeStr, err := luc.repository.GetPostLike(idUser, idPost)
	if err != nil {
		return likeStr, err
	}
	log.Println(err, likeStr)
	if !likeStr.LikeBool {
		err = luc.repository.LikePost(idPost, idUser, convertBooltoInt(like))
		log.Println(err)
		if err != nil {
			return domain.Like{}, err
		}
	} else {
		err = luc.repository.LikePostUpdate(idPost, idUser, convertBooltoInt(like))
		if err != nil {
			return domain.Like{}, err
		}
	}
	return luc.repository.GetPostLike(idPost, idUser)
}

func (luc *LikeUseCase) LikeComment(idPost, idUser int, like bool) (domain.Like, error) {
	likeStr, err := luc.repository.GetCommentLike(idUser, idPost)
	if err != nil {
		return likeStr, err
	}
	if !likeStr.LikeBool {
		err = luc.repository.LikeComment(idPost, idUser, convertBooltoInt(like))
		if err != nil {
			return domain.Like{}, err
		}
	} else {
		err = luc.repository.LikeCommentUpdate(idPost, idUser, convertBooltoInt(like))
		if err != nil {
			return domain.Like{}, err
		}
	}
	return luc.repository.GetCommentLike(idPost, idUser)
}

func convertBooltoInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
