package model

import (
	"github.com/google/uuid"
	"time"
)

type Model struct {
	Uuid   uuid.UUID `json:"uuid"`
	Date   time.Time `json:"date"`
	Title  string    `json:"title"`
	UserID string    `json:"user_id"`
}
