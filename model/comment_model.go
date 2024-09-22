package model

import "time"

type CommentRequest struct {
	Id          uint   `json:"id"`
	Description string `json:"description"`
	CardId      uint   `json:"card_id"`
	CreatedBy   string `json:"created_by"`
}

type CommentResponse struct {
	Id          uint      `json:"id"`
	Description string    `json:"description"`
	CreatedBy   string    `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UpdateCommentRequest struct {
	Description string `json:"description"`
}
