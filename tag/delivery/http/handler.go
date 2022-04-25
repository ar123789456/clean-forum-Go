package http

import (
	"encoding/json"
	"forum/models"
	"forum/tag"
	"log"
	"net/http"
)

type Handler struct {
	usecase tag.Usecase
}

func NewHandler(usecase tag.Usecase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}

type taginput struct {
	Title string `json:"title"`
}

type tagoutput struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type tagsoutput struct {
	Tags []tagoutput `json:"tags"`
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	log.Println("hello", r.Method)
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var p taginput
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.usecase.Create(p.Title)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	tags, err := h.usecase.Get()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	jtags := toTags(tags)
	result, err := json.Marshal(&jtags)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "aplication/json")
	w.Write(result)
}

func toTag(tag models.Tag) tagoutput {
	return tagoutput{
		ID:    tag.ID,
		Title: tag.Title,
	}
}

func toTags(tags []models.Tag) tagsoutput {
	var t tagsoutput
	for _, v := range tags {
		t.Tags = append(t.Tags, toTag(v))
	}
	return t
}
