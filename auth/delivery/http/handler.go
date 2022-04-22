package http

import (
	"encoding/json"
	"forum/auth"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Handler struct {
	usecase auth.UserUsecase
}

func NewHandler(usecase auth.UserUsecase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}

type SignIn struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var p SignIn
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, id, err := h.usecase.SignIn(p.Name, p.Email, p.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	cookie := http.Cookie{
		Name:   "Session_token",
		Value:  token,
		MaxAge: 300,
	}
	cookieid := http.Cookie{
		Name:  "user_id",
		Value: strconv.Itoa(id),
		// Secure: true,
	}
	http.SetCookie(w, &cookie)
	http.SetCookie(w, &cookieid)
}

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var p SignIn
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(b, &p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.usecase.SignUp(p.Name, p.Email, p.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
