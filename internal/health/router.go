package health

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/techforge-lat/linkit"
)

func SetupRoutes() {
	server := linkit.GetWithName[*echo.Echo]("server")

	server.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]any{
			"message": "Hello World!",
			"time":    time.Now(),
		})
	})
}
