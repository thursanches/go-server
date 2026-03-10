package repository

import (
	"database/sql"
	"go-server/internal/models"
)

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	// Criamos a tabela automaticamente no início
	query := `CREATE TABLE IF NOT EXISTS posts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		body TEXT
	);`
	db.Exec(query)

	return &PostRepository{db: db}
}

func (r *PostRepository) Create(p models.Post) models.Post {
	res, _ := r.db.Exec("INSERT INTO posts (body) VALUES (?)", p.Body)
	id, _ := res.LastInsertId()
	p.ID = int(id)
	return p
}

func (r *PostRepository) GetAll() []models.Post {
	rows, _ := r.db.Query("SELECT id, body FROM posts")
	defer rows.Close()

	var ps []models.Post
	for rows.Next() {
		var p models.Post
		rows.Scan(&p.ID, &p.Body)
		ps = append(ps, p)
	}
	return ps
}

func (r *PostRepository) GetByID(id int) (models.Post, bool) {
	var p models.Post
	err := r.db.QueryRow("SELECT id, body FROM posts WHERE id = ?", id).Scan(&p.ID, &p.Body)
	if err == sql.ErrNoRows {
		return p, false
	}
	return p, true
}

func (r *PostRepository) Update(id int, body string) (models.Post, bool) {
	_, err := r.db.Exec("UPDATE posts SET body = ? WHERE id = ?", body, id)
	if err != nil {
		return models.Post{}, false
	}
	return models.Post{ID: id, Body: body}, true
}

func (r *PostRepository) Delete(id int) bool {
	res, _ := r.db.Exec("DELETE FROM posts WHERE id = ?", id)
	count, _ := res.RowsAffected()
	return count > 0
}
