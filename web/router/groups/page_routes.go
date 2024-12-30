package groups

import (
	"RouteHub.Service.Dashboard/web/context"
	"RouteHub.Service.Dashboard/web/handlers"
	"github.com/labstack/echo/v4"
)

func ConfigurePageRoutes(e *echo.Echo, handlers *handlers.WebHandlers) {
	mainGroup := e.Group("/")
	mainGroup.Use(context.UserContextMiddleware(handlers.Authorizer),
		context.OrganizationContextMiddleware(handlers.Ent))

	handlers.DomainHandlers.ConfigureRoutes(mainGroup)
	handlers.SettingsHandlers.ConfigureRoutes(mainGroup)

	configureHubRoutes(mainGroup, e, handlers)
}

func configureHubRoutes(mainGroup *echo.Group, e *echo.Echo, handlers *handlers.WebHandlers) {
	hubHandlers := handlers.HubHandlers

	mainGroup.GET("", hubHandlers.HubsHandler)
	mainGroup.GET("hubs", hubHandlers.HubsHandler)
	mainGroup.GET("hubs/attach", hubHandlers.ComponentHandlers.AttachHubGet)
	mainGroup.POST("hubs/attach", hubHandlers.ComponentHandlers.AttachHubPost)

	hubGroup := e.Group("/hub/:slug")
	hubGroup.Use(context.UserContextMiddleware(handlers.Authorizer),
		context.OrganizationContextMiddleware(handlers.Ent),
		context.HubContextMiddleware(handlers.Ent))

	hubGroup.GET("", hubHandlers.IndexHandler)

	hubHandlers.CustomizeHandlers.ConfigureRoutes(hubGroup)
	hubHandlers.PagesHandlers.ConfigureRoutes(hubGroup)

	handlers.LinkHandlers.ConfigureRoutes(hubGroup)
}
