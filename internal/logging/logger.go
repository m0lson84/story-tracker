package logging

import (
	"go.uber.org/zap"
)

// Global logging instance for the application.
var logger = createLogger()

// New Creates a new sugared logger
func New() *zap.SugaredLogger {
	return logger.Sugar()
}

// NewNamed Creates a new sugared logger with the given name
func NewNamed(name string) *zap.SugaredLogger {
	return New().Named(name)
}

// createLogger Creates a new zap logger.
func createLogger() *zap.Logger {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	return logger.Named("App")
}
