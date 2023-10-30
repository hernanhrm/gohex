package config

import (
	"github.com/techforge-lat/linkit"
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

	linkit.Set[Logger](linkit.WithName("logger"), linkit.WithValue(sugarLogger))

	return sugarLogger
}
