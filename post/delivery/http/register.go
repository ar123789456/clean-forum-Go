package http

import (
	"database/sql"
	"forum/post/repository/sqlite"
	"forum/post/usecase"
	"net/http"

	//midleware
	m "forum/auth/delivery/http"
)

func RegisterPost(db *sql.DB, mux *http.ServeMux, mid m.Authentication) {
	repo := sqlite.NewPostRepository(db)
	usecase := usecase.NewPostUseCase(repo)
	handler := NewHandler(usecase)
	mux.HandleFunc("/", handler.GetAll)
	mux.HandleFunc("/post/", handler.Get)
	mux.HandleFunc("/post/user/", handler.GetUserPost)
	mux.Handle("/post/liked/user/", mid.Authentication(http.HandlerFunc(handler.GetByLike)))
	mux.Handle("/post/new", mid.Authentication(http.HandlerFunc(handler.NewPost)))
}
