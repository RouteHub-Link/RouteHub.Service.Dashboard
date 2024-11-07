package groups

import (
	"net/http"

	"RouteHub.Service.Dashboard/ent"
	"RouteHub.Service.Dashboard/features/configuration"
	"RouteHub.Service.Dashboard/web/extensions"
	"github.com/labstack/echo/v4"
)

func ConfigureAuthRoutes(e *echo.Echo, oauthConfig *configuration.OAuthConfig, authorizer *extensions.Authorizer, client *ent.Client) {

	e.GET(oauthConfig.LoginPath, func(c echo.Context) error {
		// check if user is authenticated
		_, err := authorizer.AUTHN.IsAuthenticated(c.Request())
		if err == nil {

			//app.Logger.Info("User is authenticated", "data", data)
			//userInfo := data.GetUserInfo()

			//app.Logger.Info("User info", "data", userInfo)

			// user is authenticated, redirect to home
			return c.Redirect(http.StatusFound, "/")
		}
		authorizer.AUTHN.Authenticate(c.Response().Writer, c.Request(), c.Request().RequestURI)
		return nil
	})

	e.GET(oauthConfig.CallbackPath, func(c echo.Context) error {
		authorizer.AUTHN.Callback(c.Response().Writer, c.Request())

		return nil
	})

	e.GET(oauthConfig.LogoutPath, func(c echo.Context) error {
		authorizer.AUTHN.Logout(c.Response().Writer, c.Request())
		return nil
	})
}
