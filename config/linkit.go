package config

import (
	"gohex/internal/users"

	"github.com/techforge-lat/linkit"
)

func SetupLinkit() {
	users.SetDependencies()

	// must be at the end, after every root dependency has been set
	linkit.SetAuxiliaryDependencies()
}
