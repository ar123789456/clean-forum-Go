package http

import (
	"encoding/json"
	"forum/domain"
	"forum/services/comment"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	usecase comment.CommentUseCase
}

type inputComment struct {
	Postid int    `json:"post_id"`
	Text   string `json:"text"`
}

type outputComment struct {
	ID       int    `json:"id"`
	UserID   int    `json:"user_id"`
	PostId   int    `json:"post_id"`
	UserName string `json:"username"`
	Text     string `json:"text"`
	Creat_at string `json:"create_at"`
}

type allOutputComment struct {
	Allpost []outputComment `json:"allcomments"`
}

func NewHandler(usecase comment.CommentUseCase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}

// todo userid
func (h *Handler) NewComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}
	var p inputComment
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
	comments, err := h.usecase.Create(domain.Comment{
		UserID: id,
		PostId: p.Postid,
		Text:   p.Text,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	result, err := commentsMarshal(comments)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "aplication/json")
	_, err = w.Write(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	list := strings.Split(r.URL.String(), "/")
	if len(list) == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	id, err := strconv.Atoi(list[len(list)-1])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	comments, err := h.usecase.GetAll(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	result, err := commentsMarshal(comments)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "aplication/json")
	_, err = w.Write(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

type inputDelete struct {
	Postid    int `json:"post_id"`
	CommentId int `json:"comment_id"`
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var p inputDelete
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	comments, err := h.usecase.Delete(p.CommentId, p.Postid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	result, err := commentsMarshal(comments)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "aplication/json")
	_, err = w.Write(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func commentsMarshal(allcomment []domain.Comment) ([]byte, error) {
	jsonComments := toComments(allcomment)
	return json.Marshal(&jsonComments)
}

func toComment(v domain.Comment) outputComment {
	return outputComment{
		ID:       v.ID,
		UserID:   v.UserID,
		PostId:   v.PostId,
		UserName: v.UserName,
		Text:     v.Text,
		Creat_at: v.Creat_at,
	}
}

func toComments(allcomment []domain.Comment) allOutputComment {
	var jsonAllComment allOutputComment
	for _, v := range allcomment {
		jsonAllComment.Allpost = append(jsonAllComment.Allpost, toComment(v))
	}
	return jsonAllComment
}
