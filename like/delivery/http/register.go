package http

import (
	"database/sql"
	"forum/like/repository/sqlite"
	"forum/like/usecase"
	"net/http"
)

func RegisterLike(db *sql.DB, mux *http.ServeMux) {
	repo := sqlite.NewLikeRepository(db)
	usecase := usecase.NewLikeUseCase(repo)
	handler := NewHandler(usecase)
	mux.HandleFunc("/like/post/", handler.GetPostLike)
	mux.HandleFunc("/like/comment/", handler.GetCommentLike)
	mux.HandleFunc("/like/post/add/", handler.LikePost)
	mux.HandleFunc("/like/comment/add/", handler.LikeComment)
	mux.HandleFunc("/dislike/post/add/", handler.DislikePost)
	mux.HandleFunc("/dislike/comment/add/", handler.DislikeComment)
}
