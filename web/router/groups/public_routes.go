package groups

import (
	"RouteHub.Service.Dashboard/web/handlers"
	"github.com/labstack/echo/v4"
)

func ConfigurePageRoutes(e *echo.Echo, handlers *handlers.WebHandlers) {
	pageHandlers := handlers.PageHandlers

	e.GET("/", pageHandlers.HomeHandler)
	e.GET("/hubs", pageHandlers.HubsHandler)
	e.GET("/accounts", pageHandlers.AccountsHandler)
	e.GET("/domains", pageHandlers.DomainsHandler)
}
