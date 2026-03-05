package handlers

import (
	"encoding/json"
	"go-server/internal/models"
	"go-server/internal/repository"
	"net/http"
	"strconv"
)

type PostHandler struct {
	Repo *repository.PostRepository
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var p models.Post
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	createdPost := h.Repo.Create(p)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdPost)
}

func (h *PostHandler) ListPosts(w http.ResponseWriter, r *http.Request) {
	posts := h.Repo.GetAll()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func (h *PostHandler) GetPost(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(r.URL.Path[len("/posts/"):])

	post, ok := h.Repo.GetByID(id)
	if !ok {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(post)
}

func (h *PostHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Path[len("/posts/"):])

	var input models.Post
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	post, ok := h.Repo.Update(id, input.Body)
	if !ok {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(post)
}

func (h *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Path[len("/posts/"):])
	if !h.Repo.Delete(id) {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
