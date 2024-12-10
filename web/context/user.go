package context

import (
	"net/http"

	"RouteHub.Service.Dashboard/web/extensions"
	"github.com/labstack/echo/v4"
	"github.com/zitadel/oidc/v3/pkg/oidc"
)

const (
	userContextKey = "user"
)

func GetUserFromContext(c echo.Context) (*oidc.UserInfo, error) {
	user, ok := c.Get(userContextKey).(*oidc.UserInfo)
	if !ok {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "user not found in context")
	}

	return user, nil
}

func UserContextMiddleware(authorizer *extensions.Authorizer) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user, err := getUserFromAUTHN(c, authorizer)
			if err != nil {
				c.Logger().Warn("getUserFromSlug throw an err", err)
			}

			c.Set(userContextKey, user)
			return next(c)
		}
	}
}

func getUserFromAUTHN(c echo.Context, authorizer *extensions.Authorizer) (*oidc.UserInfo, error) {
	var userInfo *oidc.UserInfo
	data, err := authorizer.AUTHN.IsAuthenticated(c.Request())

	if err != nil {
		return nil, err
	}

	userInfo = data.GetUserInfo()

	return userInfo, nil
}
