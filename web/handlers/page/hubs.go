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
	_, err := ph.GetUserInfo(c)

	if err != nil {
		return err
	}

	firstOrganization, err := ph.Ent.Organization.Query().
		First(c.Request().Context())

	ph.Logger.Info("First Organization", "data", firstOrganization)

	if err != nil {
		return err
	}

	userInfo, _ := ph.GetUserInfo(c)
	return extensions.Render(c, http.StatusOK, pages.Hubs(userInfo))
}
