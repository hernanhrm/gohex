package config

import (
	"github.com/techforge-lat/dependor"
	"go.uber.org/zap"
)

type Logger interface {
	Infow(msg string, keyAndValues ...any)
	Debugw(msg string, keyAndValues ...any)
	Warnw(msg string, keyAndValues ...any)
	Errorw(msg string, keyAndValues ...any)
}

func SetupLogger() Logger {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugarLogger := logger.Sugar()

	dependor.Set[Logger](dependor.Config{
		DependencyName: "logger",
		Value:          sugarLogger,
	})

	return sugarLogger
}
