package http

import (
	"forum/auth"
	"net/http"
)

type Handler struct {
	usecase auth.UserUsecase
}

func NewHandler(usecase auth.UserUsecase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {

}
