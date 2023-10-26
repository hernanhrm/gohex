package http

import (
	"gohex/config/dependor"

	"gohex/internal/users/application"

	"github.com/labstack/echo/v4"
)

func NewRouter() {
	controller := New(dependor.Get[application.User]("user"))

	group := dependor.Get[*echo.Echo]("echo").Group("/api/v1/users")

	group.POST("", controller.Create)
	group.PUT("", controller.Update)
	group.DELETE("", controller.Delete)
	group.GET("", controller.GetAll)
	group.DELETE("/:id", controller.GetByID)
}
