package users

import (
	"gohex/internal/users/application"
	"gohex/internal/users/infrastructure/database"

	"github.com/jackc/pgx/v5/pgxpool"
)

func BuildUseCase(db *pgxpool.Pool) application.User {
	storage := database.NewPsql(db)
	return application.New(storage)
}
