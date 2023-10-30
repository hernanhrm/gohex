package users

import (
	"gohex/internal/users/application"
	"gohex/internal/users/infrastructure/database"
	"gohex/internal/users/infrastructure/http"

	"github.com/techforge-lat/dependor"
)

func SetDependencies() {
	dependor.Set[*database.Psql](
		dependor.Config{
			AuxiliaryDependencies: map[string]string{
				"DB": "db",
			},
		},
	)

	dependor.Set[*application.User](
		dependor.Config{
			AuxiliaryDependencies: map[string]string{
				"Database": dependor.Name(database.Psql{}),
			},
		},
	)

	dependor.Set[*http.Controller](
		dependor.Config{
			AuxiliaryDependencies: map[string]string{
				"UseCase": dependor.Name(application.User{}),
			},
		},
	)
}
