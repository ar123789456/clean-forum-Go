package http

import (
	"database/sql"
	m "forum/services/auth/delivery/http"
	"forum/services/like/repository/sqlite"
	"forum/services/like/usecase"
	"net/http"
)

func RegisterLike(db *sql.DB, mux *http.ServeMux, mid m.Authentication) {
	repo := sqlite.NewLikeRepository(db)
	usecase := usecase.NewLikeUseCase(repo)
	handler := NewHandler(usecase)
	mux.HandleFunc("/like/post/", handler.GetPostLike)
	mux.HandleFunc("/like/comment/", handler.GetCommentLike)
	mux.Handle("/like/post/add/", mid.Authentication(http.HandlerFunc(handler.LikePost)))
	mux.Handle("/like/comment/add/", mid.Authentication(http.HandlerFunc(handler.LikeComment)))
	mux.Handle("/dislike/post/add/", mid.Authentication(http.HandlerFunc(handler.DislikePost)))
	mux.Handle("/dislike/comment/add/", mid.Authentication(http.HandlerFunc(handler.DislikeComment)))
}
