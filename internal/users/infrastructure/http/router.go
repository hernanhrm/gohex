package http

import (
	"gohex/config"

	"gohex/internal/users/application"
	"gohex/internal/users/infrastructure/database"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewRouter(config config.Router) {
	controller := New(BuildUseCase(config.DBPool))

	group := config.EchoHttp.Group("/api/v1/users")

	group.POST("/", controller.Create)
}

func BuildUseCase(db *pgxpool.Pool) application.User {
	storage := database.NewPsql(db)
	return application.New(storage)
}
