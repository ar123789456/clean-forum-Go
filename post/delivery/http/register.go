package http

import (
	"database/sql"
	"forum/post/repository/sqlite"
	"forum/post/usecase"
	"net/http"
)

func RegisterPost(db *sql.DB, mux *http.ServeMux) {
	repo := sqlite.NewPostRepository(db)
	usecase := usecase.NewPostUseCase(repo)
	handler := NewHandler(usecase)
	mux.HandleFunc("/", handler.GetAll)
	mux.HandleFunc("/post/", handler.Get)
	mux.HandleFunc("/post/liked/user/", handler.GetByLike)
	mux.HandleFunc("/post/user/", handler.GetUserPost)
	mux.HandleFunc("/post/new", handler.NewPost)
}
