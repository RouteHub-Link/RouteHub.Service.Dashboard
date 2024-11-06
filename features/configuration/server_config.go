package configuration

import "github.com/caarlos0/env"

type ServerConfig struct {
	Scheme string `env:"SCHEME"`
	Host   string `env:"HOST"`
	Port   string `env:"PORT"`
}

func (s *ServerConfig) IsEmpty() bool {
	if s == nil {
		return true
	}

	return s.Host == "" && s.Port == "" && s.Scheme == ""
}

func (s *ServerConfig) GetListenAddress() string {
	return s.Host + ":" + s.Port
}

func (s *ServerConfig) GetServerURL() string {
	return s.Scheme + "://" + s.Host + ":" + s.Port
}

func (s *ServerConfig) Parse(config *Config) {
	if err := env.Parse(s); err != nil {
		logger.Warn("Failed to parse server configuration from enviroment", "error", err)
	}
	if !s.IsEmpty() {
		config.Server = s
	}
}
