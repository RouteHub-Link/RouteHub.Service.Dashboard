package groups

import (
	pageHandlers "RouteHub.Service.Dashboard/web/router/handlers/page"
	"github.com/labstack/echo/v4"
)

func ConfigurePageRoutes(e *echo.Echo) {
	pageHandlers := pageHandlers.NewPageHandler()

	e.GET("/", pageHandlers.HomeHandler)
	e.GET("/hubs", pageHandlers.HubsHandler)
	e.GET("/accounts", pageHandlers.AccountsHandler)
	e.GET("/domains", pageHandlers.DomainsHandler)
}
