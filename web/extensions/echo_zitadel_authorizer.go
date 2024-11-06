package extensions

import (
	"context"
	"errors"

	"RouteHub.Service.Dashboard/features/auth"
	"github.com/labstack/echo/v4"
	"github.com/zitadel/oidc/v3/pkg/oidc"
	openid "github.com/zitadel/zitadel-go/v3/pkg/authentication/oidc"

	"github.com/zitadel/zitadel-go/v3/pkg/authentication"
)

type Authorizer struct {
	Interceptor *authentication.Interceptor[*openid.UserInfoContext[*oidc.IDTokenClaims, *oidc.UserInfo]]
	AUTHN       *authentication.Authenticator[*openid.UserInfoContext[*oidc.IDTokenClaims, *oidc.UserInfo]]
}

type Interceptor[T authentication.Ctx] struct {
	authenticator *authentication.Authenticator[T]
}

func AuthenticatorMiddleware[T authentication.Ctx](authenticator *authentication.Authenticator[T]) *Interceptor[T] {
	return &Interceptor[T]{
		authenticator: authenticator,
	}
}

// RequireAuthentication will check if there is a valid session and provide it in the context.
// If there is no session, it will automatically start a new authentication (by redirecting the user to the Login UI)
func (i *Interceptor[T]) RequireEchoAuthentication() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			ctx, err := i.authenticator.IsAuthenticated(req)
			if err != nil {
				i.authenticator.Authenticate(c.Response().Writer, req, req.RequestURI)
				return next(c)
			}

			req = req.WithContext(authentication.WithAuthContext(req.Context(), ctx))
			return next(c)
		}
	}
}

// CheckAuthentication will check if there is a valid session and provide it in the context.
// Unlike [RequireAuthentication] it will not start a new authentication if there is none.
func (i *Interceptor[T]) CheckEchoAuthentication() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			ctx, err := i.authenticator.IsAuthenticated(req)
			if err == nil {
				req = req.WithContext(authentication.WithAuthContext(req.Context(), ctx))
			}
			return next(c)
		}
	}
}

func (i *Interceptor[T]) Context(ctx context.Context) T {
	return authentication.Context[T](ctx)
}

func NewAuthorizer(oauthInstance *auth.OAuthInstance) (*Authorizer, error) {
	ctx := context.Background()
	oauthConfig := oauthInstance.GetOAuthConfig()

	authN, err := authentication.New(
		ctx,
		oauthInstance.GetZitadelInstance(),
		oauthInstance.GetSymmetricEncryptionKey(),
		openid.DefaultAuthentication(oauthConfig.ClientID, oauthInstance.GetRedirectURL(), oauthInstance.GetSymmetricEncryptionKey()),
	)

	if err != nil {
		return nil, errors.New("failed to create new authenticator")
	}

	authenticationInterceptor := authentication.Middleware(authN)

	return &Authorizer{
		Interceptor: authenticationInterceptor,
		AUTHN:       authN,
	}, nil
}
