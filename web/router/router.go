package router

import (
	"RouteHub.Service.Dashboard/features/configuration"
	"RouteHub.Service.Dashboard/web/extensions"
	"RouteHub.Service.Dashboard/web/router/groups"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ConfigureRoutes(e *echo.Echo, config *configuration.Config, authorizer *extensions.Authorizer) {
	e.Static("/static", "./web/public")
	e.Static("/static/dist/preline", "./node_modules/preline/dist")

	e.Pre(middleware.RemoveTrailingSlash())

	groups.ConfigurePageRoutes(e)
	groups.MapMiscRoutes(e)

	groups.ConfigureAuthRoutes(e, config.OAuth, authorizer)
}
