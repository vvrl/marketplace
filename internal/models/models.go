package models

import "time"

type User struct {
	ID           int64     `json:"id"`
	Login        string    `json:"login"`
	HashPassword string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
}

type Advertisement struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Text      string    `json:"text"`
	ImageURL  string    `json:"image_url"`
	Price     float64   `json:"price"`
	AuthorID  int64     `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
}
