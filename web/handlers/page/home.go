package page

import (
	"net/http"

	"RouteHub.Service.Dashboard/web/extensions"
	"RouteHub.Service.Dashboard/web/templates/pages"
	"github.com/zitadel/oidc/v3/pkg/oidc"

	"github.com/labstack/echo/v4"
)

func (ph PageHandler) HomeHandler(c echo.Context) error {
	var userInfo *oidc.UserInfo
	data, err := ph.Authorizer.AUTHN.IsAuthenticated(c.Request())

	if err == nil {
		userInfo = data.GetUserInfo()
	}

	return extensions.Render(c, http.StatusOK, pages.Home(userInfo))
}
