package http

import (
	"database/sql"
	m "forum/services/auth/delivery/http"
	"forum/services/post/repository/sqlite"
	"forum/services/post/usecase"
	"net/http"
)

func RegisterPost(db *sql.DB, mux *http.ServeMux, mid m.Authentication) {
	repo := sqlite.NewPostRepository(db)
	usecase := usecase.NewPostUseCase(repo)
	handler := NewHandler(usecase)
	mux.HandleFunc("posts/", handler.GetAll)
	mux.HandleFunc("/post/", handler.Get)
	mux.HandleFunc("/post/user/", handler.GetUserPost)
	mux.Handle("/post/liked/user/", mid.Authentication(http.HandlerFunc(handler.GetByLike)))
	mux.Handle("/post/new", mid.Authentication(http.HandlerFunc(handler.NewPost)))
}
