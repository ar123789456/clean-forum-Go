package sqlite

import (
	"database/sql"
	"forum/models"
	"log"
)

type LikeRepository struct {
	db *sql.DB
}

func NewLikeRepository(db *sql.DB) *LikeRepository {
	return &LikeRepository{
		db: db,
	}
}

type likeR struct {
	userid int
	postid int
	like   int
}

func (lr *LikeRepository) GetPostLike(id, userid int) (models.Like, error) {
	log.Println("Get post-like")
	rows, err := lr.db.Query("SELECT * FROM likes_posts WHERE id_post = ?", id)
	var like models.Like
	if err == nil {
		for rows.Next() {
			var l likeR
			rows.Scan(
				&l.userid,
				&l.postid,
				&l.like,
			)
			if l.userid == userid {
				like.LikeBool = true
			}
			if l.like == 1 {
				like.Like++
			} else {
				like.Dislike++
			}
		}
	}
	return like, err
}
func (lr *LikeRepository) GetCommentLike(id, userid int) (models.Like, error) {
	rows, err := lr.db.Query("SELECT * FROM likes_comment WHERE id_post = ?", id)
	var like models.Like
	if err == nil {
		for rows.Next() {
			var l likeR
			rows.Scan(
				&l.userid,
				&l.postid,
				&l.like,
			)
			if l.userid == userid {
				like.LikeBool = true
			}
			if l.like == 1 {
				like.Like++
			} else {
				like.Dislike++
			}
		}
	}
	return like, err
}
func (lr *LikeRepository) LikePostUpdate(idPost, idUser, like int) error {
	_, err := lr.db.Exec("UPDATE likes_posts SET liked = ? WHERE id_post = ? AND id_user = ?;", like, idPost, idUser)
	return err
}
func (lr *LikeRepository) LikeCommentUpdate(idPost, idUser, like int) error {
	_, err := lr.db.Exec("UPDATE likes_comment SET liked = ? WHERE id_comment = ? AND id_user = ?", like, idPost, idUser)
	return err
}
func (lr *LikeRepository) LikePost(idPost, idUser, like int) error {
	_, err := lr.db.Exec("INSERT INTO likes_posts(id_post, id_user, liked) VALUES(?, ?, ?);", idPost, idUser, like)
	return err
}
func (lr *LikeRepository) LikeComment(idPost, idUser, like int) error {
	_, err := lr.db.Exec("INSERT INTO likes_comment(id_post, id_user, liked) VALUES(?, ?, ?);", idPost, idUser, like)
	return err
}
