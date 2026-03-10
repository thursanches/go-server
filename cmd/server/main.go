package main

import (
	"database/sql"
	"fmt"
	"go-server/internal/handlers"
	"go-server/internal/repository"
	"log"
	"net/http"

	_ "github.com/glebarez/go-sqlite" // Driver do SQLite
)

func main() {
	// O sql.Open CRIA o arquivo .db se ele não existir
	db, err := sql.Open("sqlite", "./meu_banco.db")
	if err != nil {
		log.Fatal("Erro ao abrir o banco:", err)
	}

	defer db.Close()

	// Passamos o db para o repositório
	repo := repository.NewPostRepository(db)
	h := &handlers.PostHandler{Repo: repo}

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
			return
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
