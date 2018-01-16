package models

import (
	"time"

	"github.com/google/uuid"
)

// Tag : Tag Model
type Tag struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// Category : Category Model
type Category struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// Article : Article Model
type Article struct {
	ID         uuid.UUID  `json:"id"`
	Type       string     `json:"type"`
	Title      string     `json:"title"`
	Body       string     `json:"body"`
	Tags       []Tag      `json:"tags"`
	Categories []Category `json:"categories"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}
