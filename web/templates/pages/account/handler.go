package account

import (
	"net/http"

	"RouteHub.Service.Dashboard/web/extensions"
	"RouteHub.Service.Dashboard/web/templates/pages/account/components"
	"github.com/labstack/echo/v4"
)

type Handlers struct {
	ComponentHandlers *components.Handlers
}

func NewHandlers() *Handlers {
	return &Handlers{
		ComponentHandlers: &components.Handlers{},
	}
}

func (h *Handlers) IndexHandler(c echo.Context) error {
	return extensions.Render(c, http.StatusOK, index())
}
