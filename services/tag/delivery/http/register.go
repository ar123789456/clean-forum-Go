package http

import (
	"database/sql"
	m "forum/services/auth/delivery/http"
	"forum/services/tag/repository/sqlite"
	"forum/services/tag/usecase"
	"net/http"
)

func RegisterTag(db *sql.DB, mux *http.ServeMux, mid m.Authentication) {
	repo := sqlite.NewRepository(db)
	usc := usecase.NewUsecase(repo)
	handler := NewHandler(usc)
	mux.Handle("/tag/create", mid.Authentication(http.HandlerFunc(handler.Create)))
	mux.HandleFunc("/tags", handler.Get)
}
