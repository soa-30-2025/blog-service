package models

type Like struct {
	BlogID string `json:"blog_id"`
	UserID string `json:"user_id"`
}
