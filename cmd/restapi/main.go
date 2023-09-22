package main

import (
	"context"
	"fmt"
	"gohex/config"
	users "gohex/internal/users/infrastructure/http"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.LoadLocalConfig()
	logger := config.LoadLogger()
	dbPool, err := config.LoadDatabaseConnection(cfg)
	if err != nil {
		logger.Errorw(fmt.Sprintf("restapi.config.LoadDatabaseConnection(): %v", err))
		return
	}

	e := newEcho(cfg, echo.New().DefaultHTTPErrorHandler)

	initRoutes(cfg, config.Router{
		Logger:       logger,
		DBPool:       dbPool,
		EchoHttp:     e,
		RemoteConfig: nil,
	})

	// Start server
	go func() {
		if err := e.Start(fmt.Sprintf(":%d", cfg.ServerPort)); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

}

func newEcho(conf config.LocalConfig, errorHandler echo.HTTPErrorHandler) *echo.Echo {
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

	return e
}

func initRoutes(cfg config.LocalConfig, routerConfig config.Router) {
	users.NewRouter(routerConfig)
}
