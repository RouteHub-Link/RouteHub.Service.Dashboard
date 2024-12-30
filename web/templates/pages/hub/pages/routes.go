package pages

import "github.com/labstack/echo/v4"

func (h *Handlers) ConfigureRoutes(group *echo.Group) {
	group.GET("/pages", h.IndexGetHandler)
	group.GET("/pages/new", h.NewGetHandler)
	group.POST("/pages/new", h.NewPostHandler)

	group.GET("/pages/:pageSlug", h.EditGetHandler)
	group.POST("/pages/:pageSlug", h.EditPostHandler)
	group.GET("/pages/:pageSlug/status", h.PageStatusGet)
	group.POST("/pages/:pageSlug/status", h.PageStatusPost)

}
