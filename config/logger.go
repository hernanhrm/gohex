package config

import "go.uber.org/zap"

type Logger interface {
	Infow(msg string, keyAndValues ...any)
	Debugw(msg string, keyAndValues ...any)
	Warnw(msg string, keyAndValues ...any)
	Errorw(msg string, keyAndValues ...any)
}

func LoadLogger() Logger {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	return logger.Sugar()
}
