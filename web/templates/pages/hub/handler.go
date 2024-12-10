package hub

import (
	"log/slog"
	"net/http"

	"RouteHub.Service.Dashboard/ent"
	entHub "RouteHub.Service.Dashboard/ent/hub"
	"RouteHub.Service.Dashboard/ent/organization"
	"RouteHub.Service.Dashboard/web/extensions"
	"RouteHub.Service.Dashboard/web/handlers/page"
	"RouteHub.Service.Dashboard/web/templates/pages/hub/components"
	"github.com/labstack/echo/v4"
)

type Handlers struct {
	ComponentHandlers *components.Handlers
	RequestHandler    *page.PageRequestHandler
}

func NewHandlers(RequestHandler *page.PageRequestHandler) *Handlers {
	return &Handlers{
		RequestHandler: RequestHandler,
		ComponentHandlers: &components.Handlers{
			RequestHandler: RequestHandler,
		},
	}
}

func (h Handlers) HubsHandler(c echo.Context) error {
	slog.InfoContext(c.Request().Context(), "Accessing index page", slog.String("method", "HubsHandler"))
	_, err := h.RequestHandler.GetUserInfo(c)

	if err != nil {
		return err
	}

	userInfo, userOrganization, err := h.RequestHandler.GetUserWithOrganization(c)

	if err != nil {
		return err
	}
	h.RequestHandler.Logger.Info("User Organization", "data", userOrganization)

	hubs, err := h.RequestHandler.Ent.Hub.Query().
		Where(entHub.HasOrganizationWith(organization.ID(userOrganization.ID))).
		WithDomain().
		All(c.Request().Context())

	if err != nil {
		h.RequestHandler.Logger.Error("Error fetching hubs", "error", err)
		hubs = []*ent.Hub{}
	}

	h.RequestHandler.Logger.Info("Hubs", "data", hubs)

	return extensions.Render(c, http.StatusOK, Hubs(userInfo, hubs))
}

func (h Handlers) IndexHandler(c echo.Context) error {
	user, err := h.RequestHandler.GetUserInfo(c)
	if err != nil {
		return c.Redirect(http.StatusForbidden, "/")
	}

	hub, err := h.RequestHandler.GetHubFromSlug(c, h.RequestHandler.Ent.Hub.Query().WithDomain().WithOrganization())

	if err != nil {
		h.RequestHandler.Logger.Error("Error fetching hub", "error", err)
		return c.Redirect(http.StatusFound, "/hubs")
	}

	h.RequestHandler.Logger.Info("Hub", "data", hub)

	return extensions.Render(c, http.StatusOK, index(user, hub))
}
