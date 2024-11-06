package middlewares

import (
	"log/slog"
	"net/http"

	"RouteHub.Service.Dashboard/features/configuration"
	"RouteHub.Service.Dashboard/web/extensions"
	"github.com/labstack/echo/v4"
)

func OAuthGuard(authorizer *extensions.Authorizer, oauthConfig *configuration.OAuthConfig, logger *slog.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			logger.DebugContext(c.Request().Context(), "OauthGuard")

			pathIsAuth := c.Request().URL.Path == oauthConfig.LoginPath || c.Request().URL.Path == oauthConfig.CallbackPath || c.Request().URL.Path == oauthConfig.LogoutPath
			if !pathIsAuth {
				_, err := authorizer.AUTHN.IsAuthenticated(c.Request())

				if err == nil {
					return next(c)
				}
				return c.Redirect(http.StatusFound, oauthConfig.LoginPath)
			}
			return next(c)
		}
	}
}
