package users

import (
	"gohex/internal/users/application"
	"gohex/internal/users/infrastructure/database"
	"gohex/internal/users/infrastructure/http"

	"github.com/techforge-lat/linkit"
)

func SetDependencies() {
	linkit.Set[*database.Psql](
		linkit.WithAuxiliaryDependencies(map[string]string{
			"DB": "db",
		}),
	)

	linkit.Set[*application.User](
		linkit.WithAuxiliaryDependencies(map[string]string{
			"Database": linkit.Name(database.Psql{}),
		}),
	)

	linkit.Set[*http.Controller](
		linkit.WithAuxiliaryDependencies(map[string]string{
			"UseCase": linkit.Name(application.User{}),
		}),
	)
}
