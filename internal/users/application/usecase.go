package application

import (
	"gohex/internal/users/domain"

	"github.com/google/uuid"
	"golang.org/x/net/context"
)

type User struct {
	Database Database
}

func (u User) Create(ctx context.Context, m domain.User) error {
	return u.Database.Create(ctx, m)
}

func (u User) Update(ctx context.Context, m domain.User) error {
	return u.Database.Update(ctx, m)
}

func (u User) Delete(ctx context.Context, id uuid.UUID) error {
	return u.Database.Delete(ctx, id)
}

func (u User) List(ctx context.Context) (domain.Users, error) {
	return u.Database.List(ctx)
}

func (u User) Get(ctx context.Context, id uuid.UUID) (domain.User, error) {
	return u.Database.Get(ctx, id)
}
