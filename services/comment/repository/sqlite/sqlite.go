package sqlite

import (
	"database/sql"
	"forum/domain"
	"time"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(comment domain.Comment) error {
	timeNow := time.Now().GoString()
	_, err := r.db.Exec("INSERT INTO comments(text, id_user, id_post, create_at) VALUES(?, ?, ?, ?);", comment.Text, comment.UserID, comment.PostId, timeNow)
	return err
}

func (r *Repository) GetAll(postid int) ([]domain.Comment, error) {
	rows, err := r.db.Query(`SELECT comments.id ,comments.text, comments.id_user, comments.id_post, comments.create_at, user.NicName
	FROM comments JOIN user
	ON comments.id_user = user.ID
	WHERE comments.id_post = ?`, postid)
	var posts []domain.Comment
	if err == nil {
		for rows.Next() {
			var currentPost domain.Comment
			rows.Scan(
				&currentPost.ID,
				&currentPost.Text,
				&currentPost.UserID,
				&currentPost.PostId,
				&currentPost.Creat_at,
				&currentPost.UserName,
			)
			posts = append(posts, currentPost)
		}
		return posts, err
	}
	return posts, err
}

func (r *Repository) Delete(commentId int) error {
	_, err := r.db.Exec("delete from comments where id = $1", commentId)
	return err
}
