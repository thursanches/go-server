package repository

import (
	"go-server/internal/models"
	"sync"
)

type PostRepository struct {
	posts  map[int]models.Post
	nextID int
	mu     sync.Mutex
}

func NewPostRepository() *PostRepository {
	return &PostRepository{
		posts:  make(map[int]models.Post),
		nextID: 1,
	}
}

func (r *PostRepository) Create(p models.Post) models.Post {
	r.mu.Lock()
	defer r.mu.Unlock()
	p.ID = r.nextID
	r.nextID++
	r.posts[p.ID] = p
	return p
}

func (r *PostRepository) GetAll() []models.Post {
	r.mu.Lock()
	defer r.mu.Unlock()
	ps := make([]models.Post, 0, len(r.posts))
	for _, p := range r.posts {
		ps = append(ps, p)
	}
	return ps
}

func (r *PostRepository) GetByID(id int) (models.Post, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	p, ok := r.posts[id]
	return p, ok
}

func (r *PostRepository) Delete(id int) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.posts[id]; !ok {
		return false
	}
	delete(r.posts, id)
	return true
}

func (r *PostRepository) Update(id int, body string) (models.Post, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	p, ok := r.posts[id]
	if !ok {
		return p, false
	}
	p.Body = body
	r.posts[id] = p
	return p, true
}
