package http

import (
	"database/sql"
	"forum/category/repository/sqlite"
	"forum/category/usecase"
	"net/http"
)

func RegisterCategory(db *sql.DB, mux *http.ServeMux) {
	repo := sqlite.NewRepository(db)
	usc := usecase.NewUsecase(repo)
	handler := NewHandler(usc)
	mux.HandleFunc("/category/create", handler.Create)
	mux.HandleFunc("/categories", handler.Get)
}
