package sqlite

import (
	"database/sql"
	"forum/models"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(name string) error {
	_, err := r.db.Exec("INSERT INTO tags(title) VALUES(?);", name)
	return err
}
func (r *Repository) Get() ([]models.Tag, error) {
	rows, err := r.db.Query("SELECT * FROM tags")
	var tags []models.Tag
	if err == nil {
		for rows.Next() {
			var currentTag models.Tag
			rows.Scan(
				&currentTag.ID,
				&currentTag.Title,
			)
			tags = append(tags, currentTag)
		}
	}
	return tags, err
}
