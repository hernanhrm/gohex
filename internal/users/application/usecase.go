package application

import (
	"gohex/internal/users/domain"
	"gohex/internal/users/dto"

	"github.com/google/uuid"
	"golang.org/x/net/context"
)

type User struct {
	database Database
}

func New(database Database) User {
	return User{database: database}
}

func (u User) Create(ctx context.Context, createDto dto.Create) error {
	return u.database.Create(ctx, createDto)
}

func (u User) Update(ctx context.Context, updateDto dto.Update) error {
	return u.database.Update(ctx, updateDto)
}

func (u User) Delete(ctx context.Context, id uuid.UUID) error {
	return u.database.Delete(ctx, id)
}

func (u User) GetAll(ctx context.Context) (domain.Users, error) {
	return u.database.GetAll(ctx)
}

func (u User) GetByID(ctx context.Context, id uuid.UUID) (domain.User, error) {
	return u.database.GetByID(ctx, id)
}
