package hub

import (
	"log/slog"
	"net/http"

	"RouteHub.Service.Dashboard/ent"
	entHub "RouteHub.Service.Dashboard/ent/hub"
	"RouteHub.Service.Dashboard/ent/organization"
	"RouteHub.Service.Dashboard/web/context"
	"RouteHub.Service.Dashboard/web/extensions"
	"RouteHub.Service.Dashboard/web/templates/pages/hub/components"
	"RouteHub.Service.Dashboard/web/templates/pages/hub/customize"
	"RouteHub.Service.Dashboard/web/templates/pages/hub/pages"
	"github.com/labstack/echo/v4"
)

type Handlers struct {
	Ent               *ent.Client
	Logger            *slog.Logger
	ComponentHandlers *components.Handlers
	CustomizeHandlers *customize.Handlers
	PagesHandlers     *pages.Handlers
}

func NewHandlers(logger *slog.Logger, ent *ent.Client) *Handlers {
	return &Handlers{
		Ent:    ent,
		Logger: logger,
		ComponentHandlers: &components.Handlers{
			Logger: logger,
			Ent:    ent,
		},
		CustomizeHandlers: &customize.Handlers{
			Logger: logger,
			Ent:    ent,
		},
		PagesHandlers: &pages.Handlers{
			Logger: logger,
			Ent:    ent,
		},
	}
}

func (h Handlers) HubsHandler(c echo.Context) error {
	slog.InfoContext(c.Request().Context(), "Accessing index page", slog.String("method", "HubsHandler"))

	userInfo, err := context.GetUserFromContext(c)
	if err != nil {
		return err
	}
	userOrganization, err := context.GetOrganizationFromContext(c)

	if err != nil {
		return err
	}

	h.Logger.Info("User Organization", "data", userOrganization)

	hubs, err := h.Ent.Hub.Query().
		Where(entHub.HasOrganizationWith(organization.ID(userOrganization.ID))).
		WithDomain().
		All(c.Request().Context())

	if err != nil {
		h.Logger.Error("Error fetching hubs", "error", err)
		hubs = []*ent.Hub{}
	}

	h.Logger.Info("Hubs", "data", hubs)

	return extensions.Render(c, http.StatusOK, Hubs(userInfo, hubs))
}

func (h Handlers) IndexHandler(c echo.Context) error {

	hubFromContext := c.Get("hub").(*ent.Hub)
	h.Logger.Info("Hub", "data", hubFromContext)

	user, err := context.GetUserFromContext(c)
	if err != nil {
		return c.Redirect(http.StatusForbidden, "/")
	}

	hub, err := context.GetHubFromContext(c)

	if err != nil {
		h.Logger.Error("Error fetching hub", "error", err)
		return c.Redirect(http.StatusFound, "/hubs")
	}

	h.Logger.Info("Hub", "data", hub)

	return extensions.Render(c, http.StatusOK, index(user, hub))
}
