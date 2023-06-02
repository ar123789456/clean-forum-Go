package http

import (
	"database/sql"
	m "forum/services/auth/delivery/http"
	"forum/services/comment/repository/sqlite"
	"forum/services/comment/usecase"
	"net/http"
)

func RegisterPost(db *sql.DB, mux *http.ServeMux, mid m.Authentication) {
	repo := sqlite.NewRepository(db)
	usecase := usecase.NewUsecase(repo)
	handler := NewHandler(usecase)
	mux.HandleFunc("/comment/", handler.GetAll)
	mux.Handle("/comment/new", mid.Authentication(http.HandlerFunc(handler.NewComment)))
	mux.Handle("/comment/delete", mid.Authentication(http.HandlerFunc(handler.Delete)))
}
