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
	group.GET("accounts", pageHandlers.AccountsHandler)

	configureDomainRoutes(group, pageHandlers)
	configureHubRoutes(group, pageHandlers)
}

func configureDomainRoutes(group *echo.Group, pageHandler *page.PageHandler) {
	group.GET("domains", pageHandler.DomainsHandler)
	group.GET("domains/create", pageHandler.DomainCreateGet)
	group.POST("domains/create", pageHandler.DomainCreatePost)
}

func configureHubRoutes(group *echo.Group, pageHandler *page.PageHandler) {
	group.GET("hubs", pageHandler.HubsHandler)
	group.GET("hubs/attach", pageHandler.AttachHubGet)
	group.POST("hubs/attach", pageHandler.AttachHubPost)

	group.GET("hub/:slug", pageHandler.HubHandler)
	group.GET("hub/:slug/links", pageHandler.HubLinksHandler)
	group.GET("hub/:slug/links/create", pageHandler.HubLinksCreateHandler)
	group.POST("hub/:slug/links/create", pageHandler.HubLinkCreatePostHandler)
}
