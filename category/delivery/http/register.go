package http

import (
	"database/sql"
	"forum/category/repository/sqlite"
	"forum/category/usecase"
	"net/http"

	//midleware
	m "forum/auth/delivery/http"
)

func RegisterCategory(db *sql.DB, mux *http.ServeMux, middlewar *m.Authentication) {
	repo := sqlite.NewRepository(db)
	usc := usecase.NewUsecase(repo)
	handler := NewHandler(usc)
	mux.Handle("/category/create", middlewar.Authentication(http.HandlerFunc(handler.Create)))
	mux.HandleFunc("/categories", handler.Get)
}
