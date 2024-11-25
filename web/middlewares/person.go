package middlewares

import (
	"log/slog"

	"RouteHub.Service.Dashboard/ent"
	"RouteHub.Service.Dashboard/features/person"
	"RouteHub.Service.Dashboard/web/extensions"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"github.com/zitadel/oidc/v3/pkg/oidc"
)

func PersonMiddleware(authorizer *extensions.Authorizer, logger *slog.Logger, client *ent.Client) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			logger.Debug("PersonMiddleware called")
			ctx := c.Request().Context()

			var userInfo *oidc.UserInfo
			data, err := authorizer.AUTHN.IsAuthenticated(c.Request())
			if err != nil {
				return next(c)
			}
			userInfo = data.GetUserInfo()

			_, isnil, err := person.GetCachedPerson(ctx, userInfo.Subject)
			if !isnil {
				return next(c)
			}

			if err != nil && err != redis.Nil {
				logger.ErrorContext(ctx, "Error getting cached person", "error", err)
				return next(c)
			}

			logger.InfoContext(ctx, "User is authenticated and has Person Context is nil subject", "subject", userInfo.Subject)

			_person, err := person.GetPerson(ctx, userInfo, client, logger)
			person.SetCachedPerson(ctx, _person)

			if err != nil {
				return err
			}

			return next(c)
		}
	}
}
