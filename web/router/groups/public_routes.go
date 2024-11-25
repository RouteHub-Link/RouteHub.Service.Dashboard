package groups

import (
	"RouteHub.Service.Dashboard/web/handlers"
	"RouteHub.Service.Dashboard/web/handlers/page"
	"RouteHub.Service.Dashboard/web/middlewares"
	"github.com/labstack/echo/v4"
)

func ConfigurePageRoutes(e *echo.Echo, handlers *handlers.WebHandlers) {
	pageHandlers := handlers.PageHandlers

	group := e.Group("/")
	group.Use(middlewares.PersonMiddleware(handlers.PageHandlers.Authorizer, handlers.PageHandlers.Logger, handlers.PageHandlers.Ent))

	group.GET("", pageHandlers.HomeHandler)
	group.GET("hubs", pageHandlers.HubsHandler)
	group.GET("accounts", pageHandlers.AccountsHandler)

	configureDomainRoutes(group, pageHandlers)
}

func configureDomainRoutes(group *echo.Group, pageHandler *page.PageHandler) {
	group.GET("domains", pageHandler.DomainsHandler)
	group.GET("domains/create", pageHandler.CreateDomainGet)
	group.POST("domains/create", pageHandler.CreateDomainPost)
}
