package usecase

import (
	"forum/domain"
	auth2 "forum/services/auth"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	repository auth2.UserRepository
}

func NewUserUsecase(repository auth2.UserRepository) *UserUsecase {
	return &UserUsecase{
		repository: repository,
	}
}

func (uuc *UserUsecase) SignIn(username, email, password string) (string, int, error) {
	user, err := uuc.repository.Get(username, email)
	if err != nil {
		return "", 0, err
	}
	if !CheckPasswordHash(password, user.Password) {
		return "", 0, domain.ErrWrongPassword
	}
	user.UUID = uuid.NewV4().String()
	err = uuc.repository.Update(*user)
	return user.UUID, int(user.ID), err
}
func (uuc *UserUsecase) SignUp(user domain.User) error {
	hashPass, err := HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashPass
	return uuc.repository.Create(user)
}
func (uuc *UserUsecase) GetByToken(uuid string) (*domain.User, error) {
	return uuc.repository.GetByToken(uuid)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
