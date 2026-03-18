package repository

import (
	"database/sql"
	"go-server/internal/models"
)

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	query := `CREATE TABLE IF NOT EXISTS posts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		body TEXT,
		published BOOLEAN DEFAULT 0
	);`
	db.Exec(query)
	return &PostRepository{db: db}
}

func (r *PostRepository) Create(p models.Post) models.Post {
	res, _ := r.db.Exec("INSERT INTO posts (body, published) VALUES (?, ?)", p.Body, p.Published)
	id, _ := res.LastInsertId()
	p.ID = int(id)
	return p
}

func (r *PostRepository) GetAll() []models.Post {
	rows, err := r.db.Query("SELECT id, body, published FROM posts")
	if err != nil {
		return []models.Post{}
	}
	defer rows.Close()

	var ps []models.Post = []models.Post{}
	for rows.Next() {
		var p models.Post
		var published int
		rows.Scan(&p.ID, &p.Body, &published)
		p.Published = published == 1
		ps = append(ps, p)
	}
	return ps
}

func (r *PostRepository) GetAllPublished() []models.Post {
	rows, err := r.db.Query("SELECT id, body, published FROM posts WHERE published = 1")
	if err != nil {
		return []models.Post{}
	}
	defer rows.Close()

	var ps []models.Post = []models.Post{}
	for rows.Next() {
		var p models.Post
		var published int
		rows.Scan(&p.ID, &p.Body, &published)
		p.Published = published == 1
		ps = append(ps, p)
	}
	return ps
}

func (r *PostRepository) GetByID(id int) (models.Post, bool) {
	var p models.Post
	var published int
	err := r.db.QueryRow("SELECT id, body, published FROM posts WHERE id = ?", id).
		Scan(&p.ID, &p.Body, &published)
	if err == sql.ErrNoRows {
		return p, false
	}
	p.Published = published == 1
	return p, true
}

func (r *PostRepository) Update(id int, body string, published bool) (models.Post, bool) {
	res, err := r.db.Exec("UPDATE posts SET body = ?, published = ? WHERE id = ?", body, published, id)
	if err != nil {
		return models.Post{}, false
	}
	count, _ := res.RowsAffected()
	if count == 0 {
		return models.Post{}, false
	}
	return models.Post{ID: id, Body: body, Published: published}, true
}

func (r *PostRepository) Delete(id int) bool {
	res, _ := r.db.Exec("DELETE FROM posts WHERE id = ?", id)
	count, _ := res.RowsAffected()
	return count > 0
}
