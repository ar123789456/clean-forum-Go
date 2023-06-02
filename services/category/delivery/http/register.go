package http

import (
	"database/sql"
	m "forum/services/auth/delivery/http"
	"forum/services/category/repository/sqlite"
	"forum/services/category/usecase"
	"net/http"
)

func RegisterCategory(db *sql.DB, mux *http.ServeMux, middlewar *m.Authentication) {
	repo := sqlite.NewRepository(db)
	usc := usecase.NewUsecase(repo)
	handler := NewHandler(usc)
	mux.Handle("/category/create", middlewar.Authentication(http.HandlerFunc(handler.Create)))
	mux.HandleFunc("/categories", handler.Get)
}
