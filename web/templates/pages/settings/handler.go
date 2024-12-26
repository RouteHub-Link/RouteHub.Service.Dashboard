package settings

import (
	"log/slog"
	"net/http"

	"RouteHub.Service.Dashboard/ent"
	"RouteHub.Service.Dashboard/features/configuration"
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
	currentConfig := configuration.Get()

	return extensions.Render(c, http.StatusOK, index(userInfo, currentConfig, nil))
}
