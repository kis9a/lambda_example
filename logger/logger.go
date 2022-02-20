package logger

import "go.uber.org/zap"

var logger *zap.Logger

func Newlogger() {
	logger, _ = zap.NewProduction()
	zap.ReplaceGlobals(logger)
	defer logger.Sync()
}

func Getlogger() *zap.Logger {
	return logger
}
