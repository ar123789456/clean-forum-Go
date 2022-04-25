package http

import (
	"database/sql"
	"forum/tag/repository/sqlite"
	"forum/tag/usecase"
	"net/http"
)

func RegisterTag(db *sql.DB, mux *http.ServeMux) {
	repo := sqlite.NewRepository(db)
	usc := usecase.NewUsecase(repo)
	handler := NewHandler(usc)
	mux.HandleFunc("/tag/create", handler.Create)
	mux.HandleFunc("/tags", handler.Get)
}
