package link

import "github.com/labstack/echo/v4"

func (h *Handlers) ConfigureRoutes(group *echo.Group) {
	group.GET("/links", h.HubLinksHandler)
	group.GET("/links/create", h.HubLinksCreate)
	group.POST("/links/create", h.HubLinkCreatePost)

	group.GET("/links/:path", h.HubLinkEditGet)
	group.POST("/links/:path", h.HubLinkEditPost)

	group.GET("/links/:path/status", h.HubLinkStatusGet)
	group.POST("/links/:path/status", h.HubLinkStatusPost)
}
