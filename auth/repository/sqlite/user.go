package sqlite

import (
	"database/sql"
	"forum/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (ur *UserRepository) Create(user models.User) error {
	statement, err := ur.db.Prepare("INSERT INTO user (NicName, Email, Password) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = statement.Exec(user.Username, user.Email, user.Password)
	return err
}
func (ur *UserRepository) Get(name, email string) (*models.User, error) {
	user := models.User{}
	err := ur.db.QueryRow(
		"SELECT ID, NicName, Email, Password FROM user WHERE name=?", name).Scan(
		&user.ID, &user.Username, &user.Email, &user.Password)
	return &user, err
}
