package http

import (
	"database/sql"
	"forum/comment/repository/sqlite"
	"forum/comment/usecase"
	"net/http"
)

func RegisterPost(db *sql.DB, mux *http.ServeMux) {
	repo := sqlite.NewRepository(db)
	usecase := usecase.NewUsecase(repo)
	handler := NewHandler(usecase)
	mux.HandleFunc("/comment/", handler.GetAll)
	mux.HandleFunc("/comment/new", handler.NewComment)
	mux.HandleFunc("/comment/delete", handler.Delete)
}
