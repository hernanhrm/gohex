package bootstrap

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
)

func Run() {
	localConfig := loadLocalConfig()
	logger := setupLogger()

	_, err := setupDatabase(localConfig)
	if err != nil {
		logger.Errorw(fmt.Sprintf("restapi.config.LoadDatabaseConnection(): %v", err))
		return
	}

	server := setupEcho(localConfig, echo.New().DefaultHTTPErrorHandler)

	setupLinkit()
	setupRoutes()

	// Start server
	go func() {
		if err := server.Start(fmt.Sprintf(":%d", localConfig.ServerPort)); err != nil && !errors.Is(err, http.ErrServerClosed) {
			server.Logger.Fatalf("shutting down the server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		server.Logger.Fatal(err)
	}
}
