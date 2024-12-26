package groups

import (
	"RouteHub.Service.Dashboard/web/context"
	"RouteHub.Service.Dashboard/web/handlers"
	"RouteHub.Service.Dashboard/web/templates/pages/domain"
	"RouteHub.Service.Dashboard/web/templates/pages/hub/customize"
	"RouteHub.Service.Dashboard/web/templates/pages/hub/link"
	"RouteHub.Service.Dashboard/web/templates/pages/settings"
	"github.com/labstack/echo/v4"
)

func ConfigurePageRoutes(e *echo.Echo, handlers *handlers.WebHandlers) {
	mainGroup := e.Group("/")
	mainGroup.Use(context.UserContextMiddleware(handlers.Authorizer),
		context.OrganizationContextMiddleware(handlers.Ent))

	mainGroup.GET("", handlers.HomeHandlers.IndexHandler)
	mainGroup.GET("accounts", handlers.AccountHandlers.IndexHandler)

	configureDomainRoutes(mainGroup, handlers.DomainHandlers)
	configureSettingRoutes(mainGroup, handlers.SettingsHandlers)

	configureHubRoutes(mainGroup, e, handlers)
}

func configureDomainRoutes(group *echo.Group, handlers *domain.Handlers) {
	group.GET("domains", handlers.IndexHandler)
	group.GET("domains/create", handlers.ComponentHandlers.DomainCreateGet)
	group.POST("domains/create", handlers.ComponentHandlers.DomainCreatePost)

	group.GET("domains/partial/table", handlers.ComponentHandlers.PartialDomainTable)
}

func configureSettingRoutes(group *echo.Group, handlers *settings.Handlers) {
	group.GET("settings", handlers.IndexHandler)
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
	mapCustomize(hubGroup, customizeHandlers)

	linkHandlers := handlers.LinkHandlers
	mapLink(hubGroup, linkHandlers)
}

func mapCustomize(hubGroup *echo.Group, customizeHandlers *customize.Handlers) {
	hubGroup.GET("/customize", customizeHandlers.CustomizeGetHandler)
	hubGroup.GET("/customize/meta", customizeHandlers.MetaPageGet)
	hubGroup.GET("/customize/partial/meta", customizeHandlers.MetaPartialGet)
	hubGroup.POST("/customize/partial/meta", customizeHandlers.MetaPartialPost)

	hubGroup.GET("/customize/navbar", customizeHandlers.NavbarPageGet)
	hubGroup.GET("/customize/partial/navbar", customizeHandlers.NavbarPartialGet)
	hubGroup.POST("/customize/partial/navbar", customizeHandlers.NavbarPartialPost)

	// Two button for new navbar item and new navbar button
	hubGroup.GET("/customize/navbar/new/selection", customizeHandlers.NavbarNewSelectionGet)

	hubGroup.GET("/customize/navbar/item/:itemID/new", customizeHandlers.NavbarItemAddFormGet)
	hubGroup.POST("/customize/navbar/item/:itemID/new", customizeHandlers.NavbarItemAddFormPost)

	hubGroup.GET("/customize/navbar/item/:itemID/edit", customizeHandlers.NavbarItemEditFormGet)
	hubGroup.POST("/customize/navbar/item/:itemID/edit", customizeHandlers.NavbarItemEditFormPost)

	hubGroup.GET("/customize/navbar/item/:itemID/delete", customizeHandlers.NavbarItemDeleteFormGet)
	hubGroup.POST("/customize/navbar/item/:itemID/delete", customizeHandlers.NavbarItemDeletePost)

	hubGroup.GET("/customize/navbar/button/new", customizeHandlers.NavbarButtonEditFormGet)
	hubGroup.POST("/customize/navbar/button/new", customizeHandlers.NavbarButtonEditFormPost)

	hubGroup.GET("/customize/navbar/button/:itemID/edit", customizeHandlers.NavbarButtonEditFormGet)
	hubGroup.POST("/customize/navbar/button/:itemID/edit", customizeHandlers.NavbarButtonEditFormPost)

	hubGroup.GET("/customize/navbar/button/:itemID/delete", customizeHandlers.NavbarButtonDeleteFormGet)
	hubGroup.POST("/customize/navbar/button/:itemID/delete", customizeHandlers.NavbarButtonDeleteFormPost)

	hubGroup.GET("/customize/partial/navbar/tree", customizeHandlers.NavbarItemsTreeGet)
	hubGroup.GET("/customize/partial/navbar/shadow", customizeHandlers.NavbarItemsShadowGet)

	hubGroup.GET("/customize/footer", customizeHandlers.FooterPageGet)
	hubGroup.GET("/customize/partial/footer", customizeHandlers.FooterPartialGet)
	hubGroup.POST("/customize/partial/footer", customizeHandlers.FooterPartialPost)

	hubGroup.GET("/customize/footer/social-media-container", customizeHandlers.FooterSocialMediaContainerGet)
	hubGroup.GET("/customize/footer/social-media-links", customizeHandlers.FooterSocialMediaLinksGet)
	hubGroup.POST("/customize/footer/social-media-links", customizeHandlers.FooterSocialMediaLinksPost)
	hubGroup.GET("/customize/partial/footer/shadow", customizeHandlers.FooterShadow)

	hubGroup.GET("/customize/footer/branding-html", customizeHandlers.FooterBrandingHTMLGet)
	hubGroup.POST("/customize/footer/branding-html", customizeHandlers.FooterBrandingHTMLPost)

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
