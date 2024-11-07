package router

import (
	"log/slog"

	"RouteHub.Service.Dashboard/ent"
	"RouteHub.Service.Dashboard/features/configuration"
	"RouteHub.Service.Dashboard/web/extensions"
	"RouteHub.Service.Dashboard/web/handlers"
	"RouteHub.Service.Dashboard/web/router/groups"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ConfigureRoutes(e *echo.Echo, config *configuration.Config, logger *slog.Logger, authorizer *extensions.Authorizer, ent *ent.Client) {
	e.Static("/static", "./web/public")
	e.Static("/static/dist/preline", "./node_modules/preline/dist")

	e.Pre(middleware.RemoveTrailingSlash())

	webHandlers := handlers.NewWebHandlers(logger, authorizer, ent)

	groups.ConfigurePageRoutes(e, webHandlers)
	groups.MapMiscRoutes(e)

	groups.ConfigureAuthRoutes(e, config.OAuth, authorizer, ent)
}
