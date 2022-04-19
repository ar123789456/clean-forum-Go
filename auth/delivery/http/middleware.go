package http

import (
	"forum/auth"
	"net/http"
)

type Authentication struct {
	usecase auth.UserUsecase
}

func NewAuthentication(usecase auth.UserUsecase) *Authentication {
	return &Authentication{
		usecase: usecase,
	}
}

func (a *Authentication) Authentication(next http.Handler) {

}
