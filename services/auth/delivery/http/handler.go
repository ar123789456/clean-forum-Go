package http

import (
	"encoding/json"
	"errors"
	"forum/domain"
	"forum/services/auth"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Handler struct {
	usecase auth.UserUsecase
}

func NewHandler(usecase auth.UserUsecase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}

func (h *Handler) LoginPage(t *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if err := t.ExecuteTemplate(w, "pages/login", nil); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

type InSignIn struct {
	NameOrEmail string `json:"name_or_email"`
	Password    string `json:"password"`
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var p InSignIn
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if p.NameOrEmail == "" || p.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	token, _, err := h.usecase.SignIn(p.NameOrEmail, p.NameOrEmail, p.Password)
	if err != nil {
		if errors.Is(err, domain.ErrWrongPassword) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if errors.Is(err, domain.ErrUserNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	expiration := time.Now().Add(1000 * time.Hour)
	cookie := http.Cookie{
		Name:     "session_token",
		Value:    token,
		Expires:  expiration,
		HttpOnly: true,
		Secure:   false,
	}

	http.SetCookie(w, &cookie)
}

func (h *Handler) SignUpPage(t *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if err := t.ExecuteTemplate(w, "pages/registration", nil); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

type InSignUp struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var p InSignUp
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var user domain.User
	user.Username = p.Name
	user.Email, err = domain.NewEmail(p.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user.Password = p.Password
	err = h.usecase.SignUp(user)
	log.Println(err)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
