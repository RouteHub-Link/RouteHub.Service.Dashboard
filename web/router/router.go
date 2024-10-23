package router

import (
	"RouteHub.Service.Dashboard/web/router/groups"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ConfigureRoutes(e *echo.Echo) {
	e.Static("/static", "./web/public")
	e.Static("/static/dist/preline", "./node_modules/preline/dist")

	e.Pre(middleware.RemoveTrailingSlash())

	groups.ConfigurePageRoutes(e)
	groups.MapMiscRoutes(e)
}
