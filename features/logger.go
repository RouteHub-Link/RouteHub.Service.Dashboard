package features

import (
	"log/slog"
	"os"
	"sync"

	slogotel "github.com/remychantenay/slog-otel"
)

// LoggerConfigurer is a struct that configures the logger for the server and uses slog
var (
	logger              *slog.Logger
	onceConfigureLogger sync.Once
)

type LoggerConfigurer struct {
	Logger   *slog.Logger
	LogLevel slog.Level
}

func NewLoggerConfigurer(logLevel slog.Level) LoggerConfigurer {
	onceConfigureLogger.Do(func() {
		logger = slog.New(slogotel.OtelHandler{
			Next: slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel}),
		})

		slog.SetDefault(logger)
	})

	return LoggerConfigurer{
		Logger:   logger,
		LogLevel: logLevel,
	}
}

func GetLogger() *slog.Logger {
	return logger
}
