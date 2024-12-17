package groups

import (
	"RouteHub.Service.Dashboard/web/context"
	"RouteHub.Service.Dashboard/web/handlers"
	"RouteHub.Service.Dashboard/web/templates/pages/domain"
	"github.com/labstack/echo/v4"
)

func ConfigurePageRoutes(e *echo.Echo, handlers *handlers.WebHandlers) {
	mainGroup := e.Group("/")
	mainGroup.Use(context.UserContextMiddleware(handlers.Authorizer),
		context.OrganizationContextMiddleware(handlers.Ent))

	mainGroup.GET("", handlers.HomeHandlers.IndexHandler)
	mainGroup.GET("accounts", handlers.AccountHandlers.IndexHandler)

	configureDomainRoutes(mainGroup, handlers.DomainHandlers)
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

	mainGroup.GET("hubs", hubHandlers.HubsHandler)
	mainGroup.GET("hubs/attach", hubHandlers.ComponentHandlers.AttachHubGet)
	mainGroup.POST("hubs/attach", hubHandlers.ComponentHandlers.AttachHubPost)

	hubGroup := e.Group("/hub/:slug")
	hubGroup.Use(context.UserContextMiddleware(handlers.Authorizer),
		context.OrganizationContextMiddleware(handlers.Ent))

	hubGroup.Use(context.HubContextMiddleware(handlers.Ent))

	hubGroup.GET("", hubHandlers.IndexHandler)

	customizeHandlers := hubHandlers.CustomizeHandlers
	hubGroup.GET("/customize", customizeHandlers.CustomizeGetHandler)
	hubGroup.GET("/customize/meta", customizeHandlers.MetaGetHandler)
	hubGroup.POST("/customize/meta", customizeHandlers.MetaPostHandler)

	hubGroup.GET("/customize/navbar", customizeHandlers.NavbarGetHandler)

	hubGroup.GET("/customize/navbar/item/:itemID/new", customizeHandlers.NavbarNewItemFormGetHandler)
	hubGroup.GET("/customize/navbar/item/:itemID/edit", customizeHandlers.NavbarEditItemFormGetHandler)
	hubGroup.POST("/customize/navbar/item/:itemID/edit", customizeHandlers.NavbarItemEditFormPostHandler)

	linkHandlers := handlers.LinkHandlers
	hubGroup.GET("/links", linkHandlers.HubLinksHandler)
	hubGroup.GET("/links/create", linkHandlers.HubLinksCreateHandler)
	hubGroup.POST("/links/create", linkHandlers.HubLinkCreatePostHandler)
	hubGroup.GET("/links/:path", linkHandlers.HubLinkEditGetHandler)
	hubGroup.POST("/links/:path", linkHandlers.HubLinkEditPostHandler)
	hubGroup.GET("/links/:path/status", linkHandlers.HubLinkStatusGetHandler)
	hubGroup.POST("/links/:path/status", linkHandlers.HubLinkStatusPostHandler)
}
