package http

import (
	"database/sql"
	"forum/auth/repository/sqlite"
	"forum/auth/usecase"
	"net/http"
)

func RegisterAuth(db *sql.DB, mux *http.ServeMux) {
	repo := sqlite.NewUserRepository(db)
	usecase := usecase.NewUserUsecase(&repo)
	handler := NewHandler(usecase)
	mux.HandleFunc("/signin", handler.SignIn)
	mux.HandleFunc("/signup", handler.SignUp)
}
