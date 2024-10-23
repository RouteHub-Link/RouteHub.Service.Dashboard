package page

import (
	"net/http"

	"RouteHub.Service.Dashboard/web/extensions"
	"RouteHub.Service.Dashboard/web/templates/pages"
	"github.com/labstack/echo/v4"
)

func (ph PageHandler) HomeHandler(c echo.Context) error {
	return extensions.Render(c, http.StatusOK, pages.Home())
}
