package sqlite

import (
	"database/sql"
	"forum/domain"
	"log"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (ur *UserRepository) Create(user domain.User) error {
	statement, err := ur.db.Prepare("INSERT INTO user (NicName, Email, Password) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	log.Println(user)
	res, err := statement.Exec(user.Username, user.Email, user.Password)
	log.Println(res)
	return err
}
func (ur *UserRepository) Get(name, email string) (*domain.User, error) {
	user := domain.User{}
	err := ur.db.QueryRow(
		"SELECT ID, NicName, Email, Password FROM user WHERE NicName=? OR Email=?",
		name, email).Scan(
		&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	if user.Email == "" && user.Username == "" {
		return nil, domain.ErrUserNotFound
	}
	return &user, nil
}

func (ur *UserRepository) Update(user domain.User) error {
	statement, _ := ur.db.Prepare("UPDATE user SET Token=?  WHERE ID = ?;")
	_, err := statement.Exec(user.UUID, user.ID)
	return err
}

func (ur *UserRepository) GetByToken(uuid string) (*domain.User, error) {
	user := domain.User{}
	err := ur.db.QueryRow(
		"SELECT ID, NicName, Email, Password FROM user WHERE Token=?", uuid).Scan(
		&user.ID, &user.Username, &user.Email, &user.Password)
	if user.Email == "" && user.Username == "" {
		return nil, domain.ErrUserNotFound
	}
	return &user, err
}
