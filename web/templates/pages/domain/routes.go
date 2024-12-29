package domain

import "github.com/labstack/echo/v4"

func (h *Handlers) ConfigureRoutes(group *echo.Group) {
	group.GET("domains", h.IndexHandler)
	group.GET("domains/create", h.ComponentHandlers.DomainCreateGet)
	group.POST("domains/create", h.ComponentHandlers.DomainCreatePost)

	group.GET("domains/partial/table", h.ComponentHandlers.PartialDomainTable)
}
