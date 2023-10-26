package main

import (
	"context"
	"errors"
	"fmt"
	"gohex/config"
	"gohex/config/dependor"
	"gohex/internal"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

func main() {
	dependor.Init()

	cfg := config.LoadLocalConfig()
	dependor.Set[config.LocalConfig]("local_config", cfg)

	logger := config.LoadLogger()
	dependor.Set[config.Logger]("logger", logger)

	dbPool, err := config.LoadDatabaseConnection(cfg)
	if err != nil {
		logger.Errorw(fmt.Sprintf("restapi.config.LoadDatabaseConnection(): %v", err))
		return
	}
	dependor.Set[*pgxpool.Pool]("db", dbPool)

	e := newEcho(cfg, echo.New().DefaultHTTPErrorHandler)
	dependor.Set[*echo.Echo]("echo", e)

	internal.Load()

	initRoutes()

	// Start server
	go func() {
		if err := e.Start(fmt.Sprintf(":%d", cfg.ServerPort)); err != nil && !errors.Is(err, http.ErrServerClosed) {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with a timeout of 10 seconds.
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
