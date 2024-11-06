package configuration

import (
	"strings"

	"github.com/caarlos0/env"
)

type OAuthConfig struct {
	ClientID      string   `env:"OAUTH_CLIENTID"`
	Scopes        []string `env:"OAUTH_SCOPE"`
	AuthorizerURL string   `env:"OAUTH_AUTHORIZER_URL"`
	TokenURL      string   `env:"OAUTH_TOKEN_URL"`
	Issuer        string   `env:"OAUTH_ISSUER"`
	Domain        string   `env:"OAUTH_DOMAIN"`
	Port          string   `env:"OAUTH_PORT"`
	Insecure      bool     `env:"OAUTH_INSECURE"`
	LoginPath     string   `env:"OAUTH_LOGIN_PATH"`
	LogoutPath    string   `env:"OAUTH_LOGOUT_PATH"`
	CallbackPath  string   `env:"OAUTH_CALLBACK_PATH"`
}

func (o *OAuthConfig) IsEmpty() bool {
	if o == nil {
		return true
	}

	return o.ClientID == "" && o.CallbackPath == "" && o.AuthorizerURL == "" && o.TokenURL == "" && o.Issuer == "" && o.Domain == "" && o.Port == ""
}

func (o *OAuthConfig) GetRedirectURL(listenAddress string) string {
	redirectURL := strings.Join([]string{listenAddress, o.CallbackPath}, "")
	return redirectURL
}

func (o *OAuthConfig) Parse(config *Config) {
	if err := env.Parse(o); err != nil {
		logger.Warn("Failed to parse oauth configuration from enviroment", "error", err)
	}
	if !o.IsEmpty() {
		config.OAuth = o
	}
}
