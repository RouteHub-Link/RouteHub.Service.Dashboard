package configuration

import (
	"fmt"

	"github.com/caarlos0/env"
)

type CacheDatabaseConfig struct {
	Host         string `env:"CACHE_DATABASE_HOST"`
	Port         string `env:"CACHE_DATABASE_PORT"`
	Password     string `env:"CACHE_DATABASE_PASSWORD"`
	DB           int    `env:"CACHE_DATABASE_DB"`
	TLS          bool   `env:"CACHE_DATABASE_TLS"`
	ClearOnStart bool   `env:"CACHE_DATABASE_CLEAR_ON_START"`
}

func (r *CacheDatabaseConfig) String() string {
	return `RedisConfig; Host:` + r.Host + ` Port: ` + r.Port + ` Password: ` + r.Password + ` DB: ` + string(r.DB) + ` TLS: ` + fmt.Sprintf("%t", r.TLS) + ` ClearOnStart: ` + fmt.Sprintf("%t", r.ClearOnStart)
}

func (r *CacheDatabaseConfig) GetAddress() string {
	return r.Host + ":" + r.Port
}

func (r *CacheDatabaseConfig) IsEmpty() bool {
	if r == nil {
		return true
	}

	return r.Host == "" && r.Port == "" && r.Password == "" && r.DB == 0 && !r.TLS && !r.ClearOnStart
}

func (r *CacheDatabaseConfig) Parse(config *Config) {
	if err := env.Parse(r); err != nil {
		logger.Warn("Failed to parse cache database configuration from enviroment", "error", err)
	}
	if !r.IsEmpty() {
		config.CacheDB = r
	}
}
