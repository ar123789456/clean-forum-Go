package http

import (
	"database/sql"
	m "forum/services/auth/delivery/http"
	"forum/services/post/repository/sqlite"
	"forum/services/post/usecase"
	"html/template"
	"net/http"
)

func RegisterPost(db *sql.DB, mux *http.ServeMux, mid m.Authentication, t *template.Template) {
	repo := sqlite.NewPostRepository(db)
	usecase := usecase.NewPostUseCase(repo)
	handler := NewHandler(usecase)
	mux.HandleFunc("/posts/", handler.GetPosts(t))
	mux.HandleFunc("/post/", handler.Get)
	mux.HandleFunc("/post/user/", handler.GetUserPost)
	mux.Handle("/post/liked/user/", mid.Authentication(http.HandlerFunc(handler.GetByLike)))
	mux.Handle("/post/new", mid.Authentication(http.HandlerFunc(handler.NewPost)))
}
