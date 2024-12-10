package handlers

import (
	"log/slog"

	"RouteHub.Service.Dashboard/ent"
	"RouteHub.Service.Dashboard/web/extensions"
	"RouteHub.Service.Dashboard/web/handlers/page"
	"RouteHub.Service.Dashboard/web/templates/pages/account"
	"RouteHub.Service.Dashboard/web/templates/pages/domain"
	"RouteHub.Service.Dashboard/web/templates/pages/home"
	"RouteHub.Service.Dashboard/web/templates/pages/hub"
	"RouteHub.Service.Dashboard/web/templates/pages/hub/link"
)

type WebHandlers struct {
	PageHandlers    *page.PageRequestHandler
	AccountHandlers *account.Handlers
	DomainHandlers  *domain.Handlers
	HomeHandlers    *home.Handlers
	HubHandlers     *hub.Handlers
	LinkHandlers    *link.Handlers
}

func NewWebHandlers(logger *slog.Logger, authorizer *extensions.Authorizer, ent *ent.Client) *WebHandlers {
	pageHandlers := page.NewPageHandler(logger, authorizer, ent)

	return &WebHandlers{
		PageHandlers:    pageHandlers,
		HubHandlers:     hub.NewHandlers(pageHandlers),
		DomainHandlers:  domain.NewHandlers(pageHandlers),
		AccountHandlers: account.NewHandlers(pageHandlers),
		HomeHandlers:    home.NewHandlers(pageHandlers),
		LinkHandlers:    link.NewHandlers(pageHandlers),
	}
}
