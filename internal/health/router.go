package health

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/techforge-lat/dependor"
)

func SetupRoutes() {
	server := dependor.GetWithName[*echo.Echo]("server")

	server.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]any{
			"message": "Hello World!",
			"time":    time.Now(),
		})
	})
}
