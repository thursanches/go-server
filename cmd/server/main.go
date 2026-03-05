package main

import (
	"fmt"
	"go-server/internal/handlers"
	"go-server/internal/repository"
	"log"
	"net/http"
)

func main() {
	// Inicializa dependências
	repo := repository.NewPostRepository()
	h := &handlers.PostHandler{Repo: repo}

	// Rotas
	mux := http.NewServeMux()
	mux.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.ListPosts(w, r)
		case http.MethodPost:
			h.CreatePost(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/posts/", func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path == "/posts/" {
			http.Error(w, "ID is required", http.StatusBadRequest)
		}
		switch r.Method {
		case http.MethodGet:
			h.GetPost(w, r)
		case http.MethodPut:
			h.UpdatePost(w, r)
		case http.MethodDelete:
			h.DeletePost(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server running at :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
