package logger

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	infoLogger    *log.Logger
	debugLogger   *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
)

func init() {
	logger, _ := zap.NewProduction()
	infoLogger, _ = zap.NewStdLogAt(logger, zapcore.InfoLevel)
	debugLogger, _ = zap.NewStdLogAt(logger, zapcore.DebugLevel)
	warningLogger, _ = zap.NewStdLogAt(logger, zapcore.WarnLevel)
	errorLogger, _ = zap.NewStdLogAt(logger, zapcore.ErrorLevel)
}

func InfoLogger() *log.Logger {
	return infoLogger
}

func DebugLogger() *log.Logger {
	return debugLogger
}

func WarningLogger() *log.Logger {
	return warningLogger
}

func ErrorLogger() *log.Logger {
	return errorLogger
}
