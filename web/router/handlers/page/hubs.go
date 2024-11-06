package page

import (
	"log/slog"
	"net/http"

	"RouteHub.Service.Dashboard/web/extensions"
	"RouteHub.Service.Dashboard/web/templates/pages"
	"github.com/labstack/echo/v4"
)

func (ph PageHandler) HubsHandler(c echo.Context) error {
	slog.InfoContext(c.Request().Context(), "Accessing index page", slog.String("method", "HubsHandler"))
	return extensions.Render(c, http.StatusOK, pages.Hubs())
}
