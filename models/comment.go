package models

import "time"

type Comment struct {
	ID        string    `json:"id"`
	BlogID    string    `json:"blog_id"`
	UserID    string    `json:"user_id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
