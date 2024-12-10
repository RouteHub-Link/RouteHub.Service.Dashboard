package domain

import (
	"net/http"

	"RouteHub.Service.Dashboard/ent/person"
	"RouteHub.Service.Dashboard/web/extensions"
	"RouteHub.Service.Dashboard/web/handlers/page"
	"RouteHub.Service.Dashboard/web/templates/pages/domain/components"
	"github.com/labstack/echo/v4"
)

type Handlers struct {
	ComponentHandlers *components.Handlers
	RequestHandler    *page.PageRequestHandler
}

func NewHandlers(WebHandler *page.PageRequestHandler) *Handlers {
	return &Handlers{
		RequestHandler: WebHandler,
		ComponentHandlers: &components.Handlers{
			RequestHandler: WebHandler,
		},
	}
}

func (h *Handlers) IndexHandler(c echo.Context) error {
	userInfo, _ := h.RequestHandler.GetUserInfo(c)

	personOrganizaitonAndDomains, err := h.RequestHandler.Ent.Person.Query().
		Where(person.SubjectID(userInfo.Subject)).
		WithOrganizations().
		QueryOrganizations().
		WithDomains().
		First(c.Request().Context())

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	domains := personOrganizaitonAndDomains.Edges.Domains

	h.RequestHandler.Logger.Info("Domains", "data", domains)

	return extensions.Render(c, http.StatusOK, index(userInfo, domains))
}
