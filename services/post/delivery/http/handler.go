package http

import (
	"encoding/json"
	"forum/domain"
	"forum/services/post"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	usecase post.PostUseCase
}

func NewHandler(usecase post.PostUseCase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}

type inputPost struct {
	Title      string `json:"title"`
	CategoryID int    `json:"category"`
	Tags       []tag  `json:"tags"`
	Content    string `json:"content"`
	ID         int    `json:"id"`
	UserID     int    `json:"user_id"`
	Creat_at   string `json:"create_at"`
}

type tag struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type outputPost struct {
	Title      string `json:"title"`
	CategoryID int    `json:"category"`
	Tags       []tag  `json:"tags"`
	Content    string `json:"content"`
	ID         int    `json:"id"`
	UserID     int    `json:"user_id"`
	Creat_at   string `json:"create_at"`
}

type allOutputPost struct {
	Allpost []outputPost `json:"allpost"`
}

// todo userid
func (h *Handler) NewPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}
	var p inputPost
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
	err = h.usecase.Create(domain.Post{
		UserID:      id,
		Title:       p.Title,
		Content:     p.Content,
		CatergoryID: p.CategoryID,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	allpost, err := h.usecase.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	jsonPosts := toPosts(allpost)
	result, err := json.Marshal(&jsonPosts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(result)
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
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
	post, err := h.usecase.Get(int64(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	jsonPost := toPost(post)
	result, err := json.Marshal(&jsonPost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "aplication/json")
	w.Write(result)
}

type userLike struct {
	UserID int `json:"user_id"`
}

func (h *Handler) GetByLike(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var p userLike
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	posts, err := h.usecase.GetByLike(p.UserID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	jsonPost := toPosts(posts)
	result, err := json.Marshal(&jsonPost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "aplication/json")
	w.Write(result)
}

func (h *Handler) GetUserPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	c, err := r.Cookie("user_id")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	id, err := strconv.Atoi(c.Value)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	posts, err := h.usecase.GetUserPost(int64(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	jsonPost := toPosts(posts)
	result, err := json.Marshal(&jsonPost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "aplication/json")
	w.Write(result)
}

func toTag(tags domain.Tag) tag {
	return tag{
		Id:    tags.ID,
		Title: tags.Title,
	}
}

func toTags(tags []domain.Tag) []tag {
	var t []tag
	for _, v := range tags {
		t = append(t, toTag(v))
	}
	return t
}

func toPost(v domain.Post) outputPost {
	tags := toTags(v.Tags)
	return outputPost{
		ID:         v.ID,
		Title:      v.Title,
		CategoryID: v.CatergoryID,
		Tags:       tags,
		Content:    v.Content,
		UserID:     v.UserID,
		Creat_at:   v.Creat_at,
	}
}

func toPosts(allpost []domain.Post) allOutputPost {
	var jsonAllPost allOutputPost
	for _, v := range allpost {
		jsonAllPost.Allpost = append(jsonAllPost.Allpost, toPost(v))
	}
	return jsonAllPost
}
