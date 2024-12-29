package groups

import (
	"RouteHub.Service.Dashboard/web/context"
	"RouteHub.Service.Dashboard/web/handlers"
	"RouteHub.Service.Dashboard/web/templates/pages/domain"
	"RouteHub.Service.Dashboard/web/templates/pages/hub/link"
	"github.com/labstack/echo/v4"
)

func ConfigurePageRoutes(e *echo.Echo, handlers *handlers.WebHandlers) {
	mainGroup := e.Group("/")
	mainGroup.Use(context.UserContextMiddleware(handlers.Authorizer),
		context.OrganizationContextMiddleware(handlers.Ent))

	configureDomainRoutes(mainGroup, handlers.DomainHandlers)

	handlers.SettingsHandlers.ConfigureRoutes(mainGroup)

	configureHubRoutes(mainGroup, e, handlers)
}

func configureDomainRoutes(group *echo.Group, handlers *domain.Handlers) {
	group.GET("domains", handlers.IndexHandler)
	group.GET("domains/create", handlers.ComponentHandlers.DomainCreateGet)
	group.POST("domains/create", handlers.ComponentHandlers.DomainCreatePost)

	group.GET("domains/partial/table", handlers.ComponentHandlers.PartialDomainTable)
}

func configureHubRoutes(mainGroup *echo.Group, e *echo.Echo, handlers *handlers.WebHandlers) {
	hubHandlers := handlers.HubHandlers

	mainGroup.GET("", hubHandlers.HubsHandler)
	mainGroup.GET("hubs", hubHandlers.HubsHandler)
	mainGroup.GET("hubs/attach", hubHandlers.ComponentHandlers.AttachHubGet)
	mainGroup.POST("hubs/attach", hubHandlers.ComponentHandlers.AttachHubPost)

	hubGroup := e.Group("/hub/:slug")
	hubGroup.Use(context.UserContextMiddleware(handlers.Authorizer),
		context.OrganizationContextMiddleware(handlers.Ent))

	hubGroup.Use(context.HubContextMiddleware(handlers.Ent))

	hubGroup.GET("", hubHandlers.IndexHandler)

	hubHandlers.CustomizeHandlers.ConfigureRoutes(hubGroup)

	linkHandlers := handlers.LinkHandlers
	mapLink(hubGroup, linkHandlers)
}

func mapLink(hubGroup *echo.Group, linkHandlers *link.Handlers) {
	hubGroup.GET("/links", linkHandlers.HubLinksHandler)
	hubGroup.GET("/links/create", linkHandlers.HubLinksCreate)
	hubGroup.POST("/links/create", linkHandlers.HubLinkCreatePost)

	hubGroup.GET("/links/:path", linkHandlers.HubLinkEditGet)
	hubGroup.POST("/links/:path", linkHandlers.HubLinkEditPost)

	hubGroup.GET("/links/:path/status", linkHandlers.HubLinkStatusGet)
	hubGroup.POST("/links/:path/status", linkHandlers.HubLinkStatusPost)
}
