package config

import (
	"gohex/internal/health"
	users "gohex/internal/users/infrastructure/http"
)

func SetupRoutes() {
	health.SetupRoutes()

	users.SetupRoutes()
}
