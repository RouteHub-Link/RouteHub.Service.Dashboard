package handlers

import (
	"log/slog"

	"RouteHub.Service.Dashboard/ent"
	"RouteHub.Service.Dashboard/web/extensions"
	"RouteHub.Service.Dashboard/web/handlers/page"
)

type WebHandlers struct {
	PageHandlers *page.PageHandler
}

func NewWebHandlers(logger *slog.Logger, authorizer *extensions.Authorizer, ent *ent.Client) *WebHandlers {
	return &WebHandlers{
		PageHandlers: page.NewPageHandler(logger, authorizer, ent),
	}
}
