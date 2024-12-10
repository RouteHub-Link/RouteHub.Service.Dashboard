package home

import (
	"net/http"

	"RouteHub.Service.Dashboard/web/context"
	"RouteHub.Service.Dashboard/web/extensions"
	"github.com/labstack/echo/v4"
)

type Handlers struct{}

func NewHandlers() *Handlers {
	return &Handlers{}
}

func (h *Handlers) IndexHandler(c echo.Context) error {
	userInfo, _ := context.GetUserFromContext(c)

	return extensions.Render(c, http.StatusOK, index(userInfo))
}
