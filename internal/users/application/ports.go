package application

import (
	"context"
	"gohex/internal/users/domain"

	"github.com/google/uuid"
)

type Database interface {
	Create(ctx context.Context, m domain.User) error
	Update(ctx context.Context, m domain.User) error
	Delete(ctx context.Context, id uuid.UUID) error

	GetAll(ctx context.Context) (domain.Users, error)
	GetByID(ctx context.Context, id uuid.UUID) (domain.User, error)
}
