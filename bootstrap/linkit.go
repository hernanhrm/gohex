package bootstrap

import (
	"gohex/internal/users"

	"github.com/techforge-lat/linkit"
)

func setupLinkit() {
	users.SetDependencies()

	// must be at the end, after every root dependency has been set
	linkit.SetAuxiliaryDependencies()
}
