package configuration

import "github.com/caarlos0/env"

type DatabaseConfig struct {
	Host     string `env:"DATABASE_HOST" envDefault:"localhost"`
	Port     string `env:"DATABASE_PORT" envDefault:"5432"`
	User     string `env:"DATABASE_USER" envDefault:"postgres"`
	Password string `env:"DATABASE_PASSWORD" envDefault:"password"`
	Name     string `env:"DATABASE_NAME" envDefault:"dashboard"`
	AppName  string `env:"DATABASE_APPLICATION_NAME" envDefault:"RouteHub.Dashboard"`
	SSLMode  string `env:"DATABASE_SSL_MODE" envDefault:"disable"`
}

func (d *DatabaseConfig) String() string {
	return `DatabaseConfig; Host: ` + d.Host + ` Port: ` + d.Port + ` User: ` + d.User + ` Password: ` + d.Password + ` Name: ` + d.Name + ` AppName: ` + d.AppName + ` SSLMode: ` + d.SSLMode
}

func (d *DatabaseConfig) GetConnectionString() string {
	return "host=" + d.Host + " port=" + d.Port + " user=" + d.User + " password=" + d.Password + " dbname=" + d.Name + " application_name=" + d.AppName + " sslmode=" + d.SSLMode
}

func (d *DatabaseConfig) IsEmpty() bool {
	if d == nil {
		return true
	}

	return d.Host == "" && d.Port == "" && d.User == "" && d.Password == "" && d.Name == "" && d.AppName == "" && d.SSLMode == ""
}

func (d *DatabaseConfig) Parse(config *Config) {
	if err := env.Parse(d); err != nil {
		logger.Warn("Failed to parse database configuration from enviroment", "error", err)
	}
	if !d.IsEmpty() {
		config.Database = d
	}
}
