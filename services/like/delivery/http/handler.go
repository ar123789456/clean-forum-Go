package http

import (
	"encoding/json"
	"forum/services/like"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	usecase like.LikeUseCase
}

func NewHandler(usecase like.LikeUseCase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}

type outputLike struct {
	Like    int `json:"like"`
	Dislike int `json:"dislike"`
}

func (h *Handler) GetPostLike(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	list := strings.Split(r.URL.String(), "/")
	postId, err := strconv.Atoi(list[len(list)-1])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	c, err := r.Cookie("user_id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	userid, err := strconv.Atoi(c.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	like, err := h.usecase.GetPostLike(postId, userid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	b, err := json.Marshal(&outputLike{
		Like:    like.Like,
		Dislike: like.Dislike,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "aplication/json")
	_, err = w.Write(b)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
func (h *Handler) GetCommentLike(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	list := strings.Split(r.URL.String(), "/")
	commentId, err := strconv.Atoi(list[len(list)-1])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	c, err := r.Cookie("user_id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	userid, err := strconv.Atoi(c.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	like, err := h.usecase.GetCommentLike(commentId, userid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	b, err := json.Marshal(&outputLike{
		Like:    like.Like,
		Dislike: like.Dislike,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "aplication/json")
	_, err = w.Write(b)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
func (h *Handler) LikePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	list := strings.Split(r.URL.String(), "/")
	postId, err := strconv.Atoi(list[len(list)-1])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	c, err := r.Cookie("user_id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	id, err := strconv.Atoi(c.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	like, err := h.usecase.LikePost(postId, id, true)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	b, err := json.Marshal(&outputLike{
		Like:    like.Like,
		Dislike: like.Dislike,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "aplication/json")
	_, err = w.Write(b)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
func (h *Handler) DislikePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	list := strings.Split(r.URL.String(), "/")
	postId, err := strconv.Atoi(list[len(list)-1])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	c, err := r.Cookie("user_id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	id, err := strconv.Atoi(c.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	like, err := h.usecase.LikePost(postId, id, false)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, err := json.Marshal(&outputLike{
		Like:    like.Like,
		Dislike: like.Dislike,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "aplication/json")
	_, err = w.Write(b)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
func (h *Handler) LikeComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	list := strings.Split(r.URL.String(), "/")
	commentId, err := strconv.Atoi(list[len(list)-1])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	c, err := r.Cookie("user_id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	id, err := strconv.Atoi(c.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	like, err := h.usecase.LikeComment(commentId, id, true)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	b, err := json.Marshal(&outputLike{
		Like:    like.Like,
		Dislike: like.Dislike,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "aplication/json")
	_, err = w.Write(b)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
func (h *Handler) DislikeComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	list := strings.Split(r.URL.String(), "/")
	commentId, err := strconv.Atoi(list[len(list)-1])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	c, err := r.Cookie("user_id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	id, err := strconv.Atoi(c.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	like, err := h.usecase.LikeComment(commentId, id, false)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	b, err := json.Marshal(&outputLike{
		Like:    like.Like,
		Dislike: like.Dislike,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "aplication/json")
	_, err = w.Write(b)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
