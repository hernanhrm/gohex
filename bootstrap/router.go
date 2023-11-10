package bootstrap

import (
	"gohex/internal/health"
	users "gohex/internal/users/infrastructure/http"
)

func setupRoutes() {
	health.SetupRoutes()

	users.SetupRoutes()
}
