package sqlite

import (
	"database/sql"
	"forum/domain"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(name, description string) error {
	_, err := r.db.Exec("INSERT INTO categories(title, description) VALUES(?, ?);", name, description)
	return err
}
func (r *Repository) Get() ([]domain.Category, error) {
	rows, err := r.db.Query("SELECT * FROM categories")
	var categories []domain.Category
	if err == nil {
		for rows.Next() {
			var currentCategory domain.Category
			rows.Scan(
				&currentCategory.ID,
				&currentCategory.Title,
				&currentCategory.Description,
			)
			categories = append(categories, currentCategory)
		}
	}
	return categories, err
}
