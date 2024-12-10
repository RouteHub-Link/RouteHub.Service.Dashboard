package account

import (
	"net/http"

	"RouteHub.Service.Dashboard/web/extensions"
	"RouteHub.Service.Dashboard/web/handlers/page"
	"RouteHub.Service.Dashboard/web/templates/pages/account/components"
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

func (h *Handlers) IndexHandler(c echo.Context) error {
	return extensions.Render(c, http.StatusOK, index())
}
