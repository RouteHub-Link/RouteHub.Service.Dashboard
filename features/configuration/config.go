package configuration

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var (
	onceConfig      sync.Once
	config          *Config
	logger          *slog.Logger
	errServerConfig = errors.New("server configuration is nil")
)

type Config struct {
	Server   *ServerConfig
	OAuth    *OAuthConfig
	OTEL     *OTELConfig
	Database *DatabaseConfig
	CacheDB  *CacheDatabaseConfig
}

func (c *Config) String() string {
	return `Config; Server:` + c.Server.String() + ` OAuth: ` + c.OAuth.String() + ` OTEL: ` + c.OTEL.String() + ` Database: ` + c.Database.String()
}

type ConfigurationOptionFunc func(*Config) error

func WithLogger(l *slog.Logger) ConfigurationOptionFunc {
	return func(_ *Config) error {
		logger = l
		return nil
	}
}

func WithDefaultServerConfig() ConfigurationOptionFunc {
	return func(c *Config) error {
		c.Server = &ServerConfig{
			Host: "localhost",
			Port: "8080",
		}
		return nil
	}
}

func Get() *Config {
	return config
}

func Configure(opts ...ConfigurationOptionFunc) error {
	onceConfig.Do(func() {
		config = &Config{}

		err := godotenv.Load(".env")

		for _, opt := range opts {
			if err := opt(config); err != nil {
				return
			}
		}

		if logger == nil {
			logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
		}

		if err != nil {
			logger.Log(context.Background(), slog.LevelWarn, "Error loading .env file", slog.String("error", err.Error()))
		}

		serverConfig := &ServerConfig{}
		oauthConfig := &OAuthConfig{}
		otelConfig := &OTELConfig{}
		databaseConfig := &DatabaseConfig{}
		cacheConfig := &CacheDatabaseConfig{}

		serverConfig.Parse(config)
		oauthConfig.Parse(config)
		otelConfig.Parse(config)
		databaseConfig.Parse(config)
		cacheConfig.Parse(config)

		logger.Debug("Configuration loaded", slog.Any("config", config))
	})

	if config.Server == nil {
		logger.Error("configuration error", "error", errServerConfig)
		return errServerConfig
	}

	return nil
}
