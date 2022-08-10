package models

import (
	"time"
)

// swagger:parameters recipes newRecipe
type Recipe struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Tags        []string  `json:"tags"`
	PublishedAt time.Time `json:"publishedAt"`
}
