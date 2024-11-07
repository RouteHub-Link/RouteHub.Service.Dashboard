package middlewares

import (
	"log/slog"

	"RouteHub.Service.Dashboard/ent"
	"RouteHub.Service.Dashboard/features/person"
	"RouteHub.Service.Dashboard/web/context"
	"RouteHub.Service.Dashboard/web/extensions"
	"github.com/labstack/echo/v4"
	"github.com/zitadel/oidc/v3/pkg/oidc"
)

func PersonMiddleware(authorizer *extensions.Authorizer, logger *slog.Logger, client *ent.Client) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			logger.Debug("PersonMiddleware called")

			var userInfo *oidc.UserInfo
			data, err := authorizer.AUTHN.IsAuthenticated(c.Request())
			if err != nil {
				return next(c)
			}

			// check if the context is a ServerEchoContext
			var cc *context.ServerEchoContext

			if _, ok := c.(*context.ServerEchoContext); !ok {
				cc = &context.ServerEchoContext{c}
			}

			if cc.GetPerson() != nil {
				return next(cc)
			}

			logger.InfoContext(c.Request().Context(), "User is authenticated and has Person Context is nil")

			if err == nil {
				userInfo = data.GetUserInfo()
				person, err := person.GetPerson(userInfo, client, c.Request().Context(), logger)

				if err != nil {
					return err
				}

				cc.SetPerson(person)
			}

			return next(cc)
		}
	}
}
