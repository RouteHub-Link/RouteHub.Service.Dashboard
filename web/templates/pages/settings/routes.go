package settings

import "github.com/labstack/echo/v4"

func (h *Handlers) ConfigureRoutes(group *echo.Group) {
	group.GET("settings", h.IndexHandler)
	group.POST("settings", h.IndexPostHandler)
}
