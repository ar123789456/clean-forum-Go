package http

import (
	"database/sql"
	"forum/services/auth"
	"forum/services/auth/repository/sqlite"
	"forum/services/auth/usecase"
	"net/http"
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
		c, err := r.Cookie("session_token")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		_, err = a.usecase.GetByToken(c.Value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
