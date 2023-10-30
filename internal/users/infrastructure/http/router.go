package http

import (
	"github.com/labstack/echo/v4"
	"github.com/techforge-lat/dependor"
)

func SetupRoutes() {
	server := dependor.GetWithName[*echo.Echo]("server")
	controller := dependor.Get[*Controller]()

	group := server.Group("/api/v1/users")

	group.POST("", controller.Create)
	group.PUT("", controller.Update)
	group.DELETE("", controller.Delete)
	group.GET("", controller.List)
	group.DELETE("/:id", controller.Get)
}
