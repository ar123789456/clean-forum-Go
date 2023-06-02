package http

import (
	"database/sql"
	"forum/services/auth/repository/sqlite"
	"forum/services/auth/usecase"
	"html/template"
	"net/http"
)

func RegisterAuth(db *sql.DB, mux *http.ServeMux, t *template.Template) {
	repo := sqlite.NewUserRepository(db)
	usecase := usecase.NewUserUsecase(&repo)
	handler := NewHandler(usecase)

	//Handle pages
	mux.HandleFunc("/login", handler.LoginPage(t))
	mux.HandleFunc("/registration", handler.SignUpPage(t))

	//Handle api
	mux.HandleFunc("/api/login", handler.SignIn)
	mux.HandleFunc("/api/registration", handler.SignUp)
}
