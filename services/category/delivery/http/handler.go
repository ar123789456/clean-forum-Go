package http

import (
	"encoding/json"
	"forum/domain"
	"forum/services/category"
	"log"
	"net/http"
)

type Handler struct {
	usecase category.Usecase
}

func NewHandler(usecase category.Usecase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}

type taginput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type categoryoutput struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type categoriesoutput struct {
	Categories []categoryoutput `json:"categories"`
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
	err = h.usecase.Create(p.Title, p.Description)
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
	categories, err := h.usecase.Get()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	jcategories := toTags(categories)
	result, err := json.Marshal(&jcategories)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "aplication/json")
	w.Write(result)
}

func toTag(category domain.Category) categoryoutput {
	return categoryoutput{
		ID:          category.ID,
		Title:       category.Title,
		Description: category.Description,
	}
}

func toTags(tags []domain.Category) categoriesoutput {
	var t categoriesoutput
	for _, v := range tags {
		t.Categories = append(t.Categories, toTag(v))
	}
	return t
}
