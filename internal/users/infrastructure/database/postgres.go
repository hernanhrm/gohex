package database

import (
	"context"
	"fmt"
	"gohex/internal/users/domain"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/randallmlough/pgxscan"
)

const (
	CreateStmt  = "CreateStmt"
	UpdateStmt  = "UpdateStmt"
	DeleteStmt  = "DeleteStmt"
	GetAllStmt  = "GetAllStmt"
	GetByIDStmt = "GetByIDStmt"
)

type Psql struct {
	DB *pgxpool.Pool
}

func (p Psql) Create(ctx context.Context, m domain.User) error {
	commandTag, err := p.DB.Exec(ctx, `INSERT INTO users (id, name, email, password, created_at) VALUES ($1, $2, $3, $4, $5)`, m.ID, m.Name, m.Email, m.Password, m.CreatedAt)
	if err != nil {
		return err
	}

	if commandTag.RowsAffected() != 1 {
		return fmt.Errorf("postgres.db.Exec: %d rows affected when expecting 1", commandTag.RowsAffected())
	}

	return nil
}

func (p Psql) Update(ctx context.Context, m domain.User) error {
	commandTag, err := p.DB.Exec(ctx, `UPDATE users SET name = $1, email = $2, updated_at = $3 WHERE id = $4`, m.Name, m.Email, m.UpdatedAt, m.ID)
	if err != nil {
		return err
	}

	if commandTag.RowsAffected() != 1 {
		return fmt.Errorf("postgres.db.Exec: %d rows affected when expecting 1", commandTag.RowsAffected())
	}

	return nil
}

func (p Psql) Delete(ctx context.Context, id uuid.UUID) error {
	commandTag, err := p.DB.Exec(ctx, `DELETE FROM users WHERE id = $1`, id)
	if err != nil {
		return err
	}

	if commandTag.RowsAffected() != 1 {
		return fmt.Errorf("postgres.db.Exec: %d rows affected when expecting 1", commandTag.RowsAffected())
	}

	return nil
}

func (p Psql) List(ctx context.Context) (domain.Users, error) {
	rows, err := p.DB.Query(ctx, `SELECT id, name, email, password, created_at, updated_at FROM users`)
	if err != nil {
		return nil, fmt.Errorf("postgres.db.Query(): %w", err)
	}

	var users domain.Users
	if err := pgxscan.NewScanner(rows).Scan(&users); err != nil {
		return nil, err
	}

	return users, nil
}

func (p Psql) Get(ctx context.Context, id uuid.UUID) (domain.User, error) {
	row := p.DB.QueryRow(ctx, `SELECT id, name, email, password, created_at, updated_at FROM users WHERE id = $1`, id)

	var user domain.User
	if err := pgxscan.NewScanner(row).Scan(&user); err != nil {
		return domain.User{}, err
	}

	return user, nil
}
