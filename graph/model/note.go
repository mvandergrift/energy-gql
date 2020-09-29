package model

import (
	"time"
)

type Note struct {
	ID        int       `json:"id"`
	NoteDate  time.Time `json:"noteDate"`
	Subject   string    `json:"subject"`
	Content   string    `json:"content"`
	UserID    int
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
