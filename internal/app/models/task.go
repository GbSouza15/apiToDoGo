package models

import "github.com/google/uuid"

type Task struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserID      string    `json:"user_id,omitempty"`
}

type TaskCreate struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
