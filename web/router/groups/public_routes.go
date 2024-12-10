package groups

import (
	"RouteHub.Service.Dashboard/web/handlers"
	"RouteHub.Service.Dashboard/web/middlewares"
	"RouteHub.Service.Dashboard/web/templates/pages/domain"
	"github.com/labstack/echo/v4"
)

func ConfigurePageRoutes(e *echo.Echo, handlers *handlers.WebHandlers) {
	mainGroup := e.Group("/")
	mainGroup.Use(middlewares.PersonMiddleware(handlers.PageHandlers.Authorizer, handlers.PageHandlers.Logger, handlers.PageHandlers.Ent))

	mainGroup.GET("", handlers.HomeHandlers.IndexHandler)
	mainGroup.GET("accounts", handlers.AccountHandlers.IndexHandler)

	configureDomainRoutes(mainGroup, handlers.DomainHandlers)
	configureHubRoutes(mainGroup, e, handlers)
}

func configureDomainRoutes(group *echo.Group, handlers *domain.Handlers) {
	group.GET("domains", handlers.IndexHandler)
	group.GET("domains/create", handlers.ComponentHandlers.DomainCreateGet)
	group.POST("domains/create", handlers.ComponentHandlers.DomainCreatePost)
}

func configureHubRoutes(mainGroup *echo.Group, e *echo.Echo, handlers *handlers.WebHandlers) {
	hubHandlers := handlers.HubHandlers
	mainGroup.GET("hubs", hubHandlers.HubsHandler)
	mainGroup.GET("hubs/attach", hubHandlers.ComponentHandlers.AttachHubGet)
	mainGroup.POST("hubs/attach", hubHandlers.ComponentHandlers.AttachHubPost)

	hubGroup := e.Group("/hub/:slug")
	hubGroup.Use(middlewares.PersonMiddleware(handlers.PageHandlers.Authorizer, handlers.PageHandlers.Logger, handlers.PageHandlers.Ent))

	hubGroup.GET("", hubHandlers.IndexHandler)

	linkHandlers := handlers.LinkHandlers
	hubGroup.GET("/links", linkHandlers.HubLinksHandler)
	hubGroup.GET("/links/create", linkHandlers.HubLinksCreateHandler)
	hubGroup.POST("/links/create", linkHandlers.HubLinkCreatePostHandler)
}
