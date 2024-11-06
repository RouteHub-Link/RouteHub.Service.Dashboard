package auth

import (
	"errors"
	"log/slog"
	"os"
	"sync"

	"RouteHub.Service.Dashboard/features/configuration"
	"github.com/zitadel/zitadel-go/v3/pkg/zitadel"
	"golang.org/x/exp/rand"
)

var (
	onceConfigureZitadel sync.Once
	oauthInstance        *OAuthInstance
)

type OAuthInstance struct {
	zitadelInstance        *zitadel.Zitadel
	oauthConfig            *configuration.OAuthConfig
	logger                 *slog.Logger
	symmetricEncryptionKey string
	redirectURL            string
}

func NewOAuthInstance(opts ...OAuthInstanceOptionFunc) (instance *OAuthInstance, err error) {
	onceConfigureZitadel.Do(func() {
		err = oauthConfigure(opts)
	})

	if err != nil {
		return nil, err
	}

	if oauthInstance.oauthConfig == nil {
		return nil, errors.New("oauth config is nil")
	}

	return oauthInstance, nil
}

func oauthConfigure(opts []OAuthInstanceOptionFunc) (err error) {
	oauthInstance = &OAuthInstance{}

	for _, opt := range opts {
		if err = opt(oauthInstance); err != nil {
			return
		}
	}

	if oauthInstance.logger == nil {
		oauthInstance.logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	if err != nil {
		oauthInstance.logger.Error("new oauth instance", "error", err)
		return
	}

	if oauthInstance.symmetricEncryptionKey == "" {
		oauthInstance.SetSymmetricEncryptionKey(randomString(32))
	}

	oauthConfig := oauthInstance.GetOAuthConfig()

	if oauthConfig.IsEmpty() {
		oauthInstance.logger.Error("new oauth instance", "error", "oauth config is empty")
		return errors.New("oauth config is empty")
	}

	serverConfig := configuration.Get().Server

	oauthInstance.SetRedirectURL(oauthConfig.GetRedirectURL(serverConfig.GetServerURL()))

	if oauthInstance.GetRedirectURL() == "" {
		oauthInstance.logger.Error("new oauth instance", "error", "redirect url is empty")
		return errors.New("redirect url is empty")
	}

	if oauthConfig.Insecure {
		oauthInstance.SetZitadelInstance(zitadel.New(oauthConfig.Domain, zitadel.WithInsecure(oauthConfig.Port)))
	} else {
		oauthInstance.SetZitadelInstance(zitadel.New(oauthConfig.Domain))
	}
	return nil
}

func randomString(length int) string {
	// use crypto/rand to generate a random string
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
