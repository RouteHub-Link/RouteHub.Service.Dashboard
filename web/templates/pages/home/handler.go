package home

import (
	"net/http"

	"RouteHub.Service.Dashboard/web/extensions"
	"RouteHub.Service.Dashboard/web/handlers/page"
	"github.com/labstack/echo/v4"
	"github.com/zitadel/oidc/v3/pkg/oidc"
)

type Handlers struct {
	RequestHandler *page.PageRequestHandler
}

func NewHandlers(RequestHandler *page.PageRequestHandler) *Handlers {
	return &Handlers{
		RequestHandler: RequestHandler,
	}
}

func (h *Handlers) IndexHandler(c echo.Context) error {
	var userInfo *oidc.UserInfo
	data, err := h.RequestHandler.Authorizer.AUTHN.IsAuthenticated(c.Request())

	if err == nil {
		userInfo = data.GetUserInfo()
	}

	return extensions.Render(c, http.StatusOK, index(userInfo))
}
