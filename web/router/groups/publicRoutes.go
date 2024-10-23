package groups

import (
	"RouteHub.Service.Dashboard/web/router/handlers"
	"github.com/labstack/echo/v4"
)

func ConfigurePageRoutes(e *echo.Echo) {
	webHandlers := handlers.NewWebHandler()

	e.GET("/", webHandlers.HomeHandler)
	e.GET("/hubs", webHandlers.HubsHandler)
}
