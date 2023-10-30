package config

import (
	"gohex/internal/users"

	"github.com/techforge-lat/dependor"
)

func SetupDependor() {
	users.SetDependencies()

	// must be at the end, after every root dependency has been set
	dependor.SetAuxiliarDependencies()
}
