package http

import (
	"gohex/config"

	"gohex/internal/users"
)

func NewRouter(config config.Router) {
	controller := New(users.BuildUseCase(config.DBPool))

	group := config.EchoHttp.Group("/api/v1/users")

	group.POST("", controller.Create)
	group.PUT("", controller.Update)
	group.DELETE("", controller.Delete)
	group.GET("", controller.GetAll)
	group.DELETE("/:id", controller.GetByID)

}
