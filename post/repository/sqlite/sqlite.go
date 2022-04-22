package sqlite

import (
	"database/sql"
	"forum/models"
	"time"
)

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{
		db: db,
	}
}

func (pr *PostRepository) Create(post models.Post) error {
	statement, err := pr.db.Prepare("INSERT INTO posts(title, content, create_at, update_at, id_user) VALUES(?, ?, ?, ?, ?);")
	if err != nil {
		return err
	}
	timeNow := time.Now().GoString()
	_, err = statement.Exec(post.Title, post.Content, timeNow, timeNow, post.UserID)
	return err
}
func (pr *PostRepository) Get(id int64) (models.Post, error) {
	var post models.Post
	err := pr.db.QueryRow(
		"SELECT id, title, content, create_at, update_at, id_user FROM posts WHERE id=?", id).Scan(
		&post.ID, &post.Title, &post.Content, &post.Creat_at, &post.Update_to, &post.UserID)
	return post, err
}
func (pr *PostRepository) GetAll() ([]models.Post, error) {
	rows, err := pr.db.Query("SELECT * FROM posts")
	var posts []models.Post
	if err == nil {
		for rows.Next() {
			var currentPost models.Post
			rows.Scan(
				&currentPost.ID,
				&currentPost.Title,
				&currentPost.Content,
				&currentPost.Creat_at,
				&currentPost.Update_to,
				&currentPost.UserID,
			)
			posts = append(posts, currentPost)
		}
		return posts, err
	}
	return posts, err
}
func (pr *PostRepository) GetByLike() ([]models.Post, error) {
	return []models.Post{}, nil
}
func (pr *PostRepository) GetUserPost(id int64) ([]models.Post, error) {
	rows, err := pr.db.Query("SELECT * FROM posts WHERE id_user=?;", id)
	var posts []models.Post
	if err == nil {
		for rows.Next() {
			var currentPost models.Post
			rows.Scan(
				&currentPost.ID,
				&currentPost.Title,
				&currentPost.Content,
				&currentPost.Creat_at,
				&currentPost.Update_to,
				&currentPost.UserID,
			)
			posts = append(posts, currentPost)
		}
	}
	return posts, err
}
