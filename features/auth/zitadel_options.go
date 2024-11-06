package auth

import (
	"log/slog"

	"RouteHub.Service.Dashboard/features/configuration"
	"github.com/zitadel/zitadel-go/v3/pkg/zitadel"
)

type OAuthInstanceOptionFunc func(*OAuthInstance) error

func WithSymmetricEncryptionKey(symmetricEncryptionKey string) OAuthInstanceOptionFunc {
	return func(o *OAuthInstance) error {
		o.SetSymmetricEncryptionKey(symmetricEncryptionKey)
		return nil
	}
}

func WithZitadelInstance(zitadelInstance *zitadel.Zitadel) OAuthInstanceOptionFunc {
	return func(o *OAuthInstance) error {
		o.SetZitadelInstance(zitadelInstance)
		return nil
	}
}

func WithOAuthConfig(oauthConfig *configuration.OAuthConfig) OAuthInstanceOptionFunc {
	return func(o *OAuthInstance) error {
		o.SetOAuthConfig(oauthConfig)
		return nil
	}
}

func WithLogger(logger *slog.Logger) OAuthInstanceOptionFunc {
	return func(o *OAuthInstance) error {
		o.SetLogger(logger)
		return nil
	}
}

func WithRedirectURL(redirectURL string) OAuthInstanceOptionFunc {
	return func(o *OAuthInstance) error {
		o.SetRedirectURL(redirectURL)
		return nil
	}
}

func (OAI *OAuthInstance) SetSymmetricEncryptionKey(symmetricEncryptionKey string) {
	OAI.symmetricEncryptionKey = symmetricEncryptionKey
}

func (OAI *OAuthInstance) SetZitadelInstance(zitadelInstance *zitadel.Zitadel) {
	OAI.zitadelInstance = zitadelInstance
}

func (OAI *OAuthInstance) SetLogger(logger *slog.Logger) {
	OAI.logger = logger
}

func (OAI *OAuthInstance) SetOAuthConfig(oauthConfig *configuration.OAuthConfig) {
	OAI.oauthConfig = oauthConfig
}

func (OAI *OAuthInstance) GetOAuthConfig() *configuration.OAuthConfig {
	return OAI.oauthConfig
}

func (OAI *OAuthInstance) GetLogger() *slog.Logger {
	return OAI.logger
}

func (OAI *OAuthInstance) GetZitadelInstance() *zitadel.Zitadel {
	return OAI.zitadelInstance
}

func (OAI *OAuthInstance) GetSymmetricEncryptionKey() string {
	return OAI.symmetricEncryptionKey
}

func (OAI *OAuthInstance) GetRedirectURL() string {
	return OAI.redirectURL
}

func (OAI *OAuthInstance) SetRedirectURL(redirectURL string) {
	OAI.redirectURL = redirectURL
}
