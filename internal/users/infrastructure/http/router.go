package http

import (
	"gohex/config"

	"gohex/internal/users"
)

func NewRouter(config config.Router) {
	controller := New(users.BuildUseCase(config.DBPool))

	group := config.EchoHttp.Group("/api/v1/users")

	group.POST("/", controller.Create)
}
