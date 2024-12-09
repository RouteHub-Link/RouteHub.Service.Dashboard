package page

import (
	"net/http"

	"RouteHub.Service.Dashboard/web/extensions"
	"RouteHub.Service.Dashboard/web/templates/pages"
	"github.com/labstack/echo/v4"
)

func (ph PageHandler) HubLinksHandler(c echo.Context) error {
	user, err := ph.GetUserInfo(c)
	if err != nil {
		return c.Redirect(http.StatusForbidden, "/")
	}

	query := ph.Ent.Hub.Query().WithDomain().WithLinks()

	hub, err := ph.GetHubFromSlug(c, query)

	if err != nil {
		ph.Logger.Error("Error fetching hub", "error", err)
		return c.Redirect(http.StatusFound, "/hubs")
	}

	return extensions.Render(c, http.StatusOK, pages.HubLinks(user, hub, hub.Edges.Links))
}
