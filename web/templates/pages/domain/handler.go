package domain

import (
	"log/slog"
	"net/http"

	"RouteHub.Service.Dashboard/ent"
	"RouteHub.Service.Dashboard/ent/person"
	"RouteHub.Service.Dashboard/web/context"
	"RouteHub.Service.Dashboard/web/extensions"
	"RouteHub.Service.Dashboard/web/templates/pages/domain/components"
	"github.com/labstack/echo/v4"
)

type Handlers struct {
	Ent               *ent.Client
	Logger            *slog.Logger
	ComponentHandlers *components.Handlers
}

func NewHandlers(logger *slog.Logger, ent *ent.Client) *Handlers {
	return &Handlers{
		Ent:               ent,
		Logger:            logger,
		ComponentHandlers: components.NewHandlers(logger, ent),
	}
}

func (h *Handlers) IndexHandler(c echo.Context) error {
	userInfo, _ := context.GetUserFromContext(c)

	personOrganizaitonAndDomains, err := h.Ent.Person.Query().
		Where(person.SubjectID(userInfo.Subject)).
		WithOrganizations().
		QueryOrganizations().
		WithDomains().
		First(c.Request().Context())

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	domains := personOrganizaitonAndDomains.Edges.Domains

	h.Logger.Info("Domains", "data", domains)

	return extensions.Render(c, http.StatusOK, index(userInfo, domains))
}
