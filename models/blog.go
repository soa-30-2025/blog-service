package models

import "time"

type Blog struct {
	ID       string `json:"id"`
    Title string `json:"title"`
    Description    string `json:"description"`
    AuthorId string `json:"author_id"`
    CreatedAt   time.Time `json:"created_at"`
}