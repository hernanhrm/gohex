package users

import (
	"gohex/config/dependor"
	"gohex/internal/users/application"
	"gohex/internal/users/infrastructure/database"

	"github.com/jackc/pgx/v5/pgxpool"
)

// BuildAndLoad builds and loads `users` dependency in dependor
func BuildAndLoad() {
	storage := database.NewPsql(dependor.Get[*pgxpool.Pool]("db"))
	app := application.New(storage)

	dependor.Set[application.User]("user", app)
}
