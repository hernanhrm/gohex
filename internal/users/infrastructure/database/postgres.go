package database

import (
	"context"
	"fmt"
	"gohex/internal/users/domain"
	"gohex/internal/users/dto"

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
	db *pgxpool.Pool
}

func NewPsql(db *pgxpool.Pool) Psql {
	return Psql{db: db}
}

func (p Psql) Create(ctx context.Context, createDto dto.Create) error {
	commandTag, err := p.db.Exec(ctx, `INSERT INTO users (id, name, email, password, created_at) VALUES ($1, $2, $3, $4, $5)`, createDto.ID, createDto.Name, createDto.Email, createDto.Password, createDto.CreatedAt)
	if err != nil {
		return err
	}

	if commandTag.RowsAffected() != 1 {
		return fmt.Errorf("postgres.db.Exec: %d rows affected when expecting 1", commandTag.RowsAffected())
	}

	return nil
}

func (p Psql) Update(ctx context.Context, updateDto dto.Update) error {
	commandTag, err := p.db.Exec(ctx, `UPDATE users SET name = $1, email = $2, updated_at = $3 WHERE id = $4`, updateDto.Name, updateDto.Email, updateDto.UpdatedAt, updateDto.ID)
	if err != nil {
		return err
	}

	if commandTag.RowsAffected() != 1 {
		return fmt.Errorf("postgres.db.Exec: %d rows affected when expecting 1", commandTag.RowsAffected())
	}

	return nil
}

func (p Psql) Delete(ctx context.Context, id uuid.UUID) error {
	commandTag, err := p.db.Exec(ctx, `DELETE FROM users WHERE id = $1`, id)
	if err != nil {
		return err
	}

	if commandTag.RowsAffected() != 1 {
		return fmt.Errorf("postgres.db.Exec: %d rows affected when expecting 1", commandTag.RowsAffected())
	}

	return nil
}

func (p Psql) GetAll(ctx context.Context) (domain.Users, error) {
	rows, err := p.db.Query(ctx, `SELECT id, name, email, password, created_at, updated_at FROM users`)
	if err != nil {
		return nil, fmt.Errorf("postgres.db.Query(): %w", err)
	}

	var users dto.QueryUsers
	if err := pgxscan.NewScanner(rows).Scan(&users); err != nil {
		return nil, err
	}

	return users.AsDomainUsers(), nil
}

func (p Psql) GetByID(ctx context.Context, id uuid.UUID) (domain.User, error) {
	row := p.db.QueryRow(ctx, `SELECT id, name, email, password, created_at, updated_at FROM users WHERE id = $1`, id)

	var user dto.QueryUser
	if err := pgxscan.NewScanner(row).Scan(&user); err != nil {
		return domain.User{}, err
	}

	return user.AsDomainUser(), nil
}
