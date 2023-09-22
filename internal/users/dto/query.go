package dto

import (
	"gohex/internal/users/domain"
	"time"

	"github.com/google/uuid"
)

type QueryUser struct {
	ID        uuid.UUID `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (q QueryUser) AsDomainUser() domain.User {
	return domain.User{
		ID:        q.ID,
		Name:      q.Name,
		Email:     q.Email,
		Password:  q.Password,
		CreatedAt: q.CreatedAt,
		UpdatedAt: q.UpdatedAt,
	}
}

type QueryUsers []QueryUser

func (q QueryUsers) AsDomainUsers() domain.Users {
	users := make(domain.Users, len(q))
	for _, v := range q {
		users = append(users, v.AsDomainUser())
	}

	return users
}
