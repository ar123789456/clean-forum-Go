package usecase

import (
	"forum/auth"
	"forum/models"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	repository auth.UserRepository
}

func NewUserUsecase(repository auth.UserRepository) *UserUsecase {
	return &UserUsecase{
		repository: repository,
	}
}

func (uuc *UserUsecase) SignIn(username, email, password string) (string, error) {
	user, err := uuc.repository.Get(username, email)
	if err != nil {
		return "", err
	}
	if !CheckPasswordHash(password, user.Password) {
		return "", auth.WrongPassword
	}
	user.UUID = uuid.NewV4().String()
	err = uuc.repository.Update(*user)
	return user.UUID, err
}
func (uuc *UserUsecase) SignUp(name, email, password string) error {
	user := models.User{}
	hashPass, err := HashPassword(password)
	if err != nil {
		return err
	}
	user.Email = email
	user.Username = name
	user.Password = hashPass
	return uuc.repository.Create(user)
}
func (uuc *UserUsecase) ParseByToken(string) (*models.User, error)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
