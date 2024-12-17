package customize

import (
	"net/http"

	"RouteHub.Service.Dashboard/web/context"
	"RouteHub.Service.Dashboard/web/extensions"
	"RouteHub.Service.Dashboard/web/templates/pages/partial"
	"github.com/labstack/echo/v4"
)

func (h Handlers) NavbarItemsTreeGetHandler(c echo.Context) error {
	hub, _ := context.GetHubFromContext(c)

	extensions.HTMXAppendPrelineRefresh(c)
	return extensions.Render(c, http.StatusOK, partial.TreeView(NavbarToTree(&hub.HubDetails.NavbarDescription, hub.Slug)))
}

func (h Handlers) NavbarItemsShadowGetHandler(c echo.Context) error {
	hub, _ := context.GetHubFromContext(c)

	return extensions.Render(c, http.StatusOK, partial.ShadowHub(NavbarShadow(&hub.HubDetails.NavbarDescription)))
}

func (h Handlers) NavbarNewSelectionGetHandler(c echo.Context) error {
	hub, _ := context.GetHubFromContext(c)

	return extensions.Render(c, http.StatusOK, navbarCreateSelector(hub.Slug))
}
