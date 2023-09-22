package http

import (
	"context"
	"gohex/internal/users/domain"
	"gohex/internal/users/dto"

	"github.com/google/uuid"
)

type UseCase interface {
	Create(context.Context, dto.Create) error
	Update(context.Context, dto.Update) error
	Delete(context.Context, uuid.UUID) error

	GetAll(context.Context) (domain.Users, error)
	GetByID(ctx context.Context, id uuid.UUID) (domain.User, error)
}
