package models

type Post struct {
	ID        int    `json:"id"`
	Body      string `json:"body"`
	Published bool   `json:"published"`
}
