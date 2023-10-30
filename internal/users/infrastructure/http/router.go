package http

import (
	"github.com/labstack/echo/v4"
	"github.com/techforge-lat/linkit"
)

func SetupRoutes() {
	server := linkit.GetWithName[*echo.Echo]("server")
	controller := linkit.Get[*Controller]()

	group := server.Group("/api/v1/users")

	group.POST("", controller.Create)
	group.PUT("", controller.Update)
	group.DELETE("", controller.Delete)
	group.GET("", controller.List)
	group.DELETE("/:id", controller.Get)
}
