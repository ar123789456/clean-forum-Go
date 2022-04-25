package http

import (
	"database/sql"
	"forum/tag/repository/sqlite"
	"forum/tag/usecase"
	"net/http"

	//midleware
	m "forum/auth/delivery/http"
)

func RegisterTag(db *sql.DB, mux *http.ServeMux, mid m.Authentication) {
	repo := sqlite.NewRepository(db)
	usc := usecase.NewUsecase(repo)
	handler := NewHandler(usc)
	mux.Handle("/tag/create", mid.Authentication(http.HandlerFunc(handler.Create)))
	mux.HandleFunc("/tags", handler.Get)
}
