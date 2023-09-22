package dto

import (
	"time"

	"github.com/google/uuid"
)

type Update struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdatePassword struct {
	Password string `json:"password"`
}
