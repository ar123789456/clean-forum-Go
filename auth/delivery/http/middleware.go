package http

import (
	"database/sql"
	"forum/auth"
	"forum/auth/repository/sqlite"
	"forum/auth/usecase"
	"net/http"
	"strconv"
	"time"
)

type Authentication struct {
	usecase auth.UserUsecase
}

func NewAuthentication(db *sql.DB) *Authentication {
	repo := sqlite.NewUserRepository(db)
	usecase := usecase.NewUserUsecase(&repo)
	return &Authentication{
		usecase: usecase,
	}
}

func (a *Authentication) Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("Session_token")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		user, err := a.usecase.ParseByToken(c.Value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		c, err = r.Cookie("user_id")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if strconv.Itoa(int(user.ID)) != c.Value {
			c = &http.Cookie{
				Name:    "user_id",
				Value:   "",
				Path:    "/",
				Expires: time.Unix(0, 0),

				HttpOnly: true,
			}
			http.SetCookie(w, c)
			c = &http.Cookie{
				Name:    "Session_token",
				Value:   "",
				Path:    "/",
				Expires: time.Unix(0, 0),

				HttpOnly: true,
			}
			http.SetCookie(w, c)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
