package bootstrap

import (
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/techforge-lat/linkit"
)

func setupEcho(conf LocalConfig, errorHandler echo.HTTPErrorHandler) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.Secure())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: time.Minute,
	}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: strings.Split(conf.AllowedDomains, ","),
		AllowMethods: strings.Split(conf.AllowedMethods, ","),
	}))

	e.HTTPErrorHandler = errorHandler

	linkit.Set[*echo.Echo](linkit.WithName("server"), linkit.WithValue(e))

	return e
}
