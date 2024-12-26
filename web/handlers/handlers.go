package handlers

import (
	"log/slog"

	"RouteHub.Service.Dashboard/ent"
	"RouteHub.Service.Dashboard/web/extensions"
	"RouteHub.Service.Dashboard/web/templates/pages/account"
	"RouteHub.Service.Dashboard/web/templates/pages/domain"
	"RouteHub.Service.Dashboard/web/templates/pages/home"
	"RouteHub.Service.Dashboard/web/templates/pages/hub"
	"RouteHub.Service.Dashboard/web/templates/pages/hub/link"
	"RouteHub.Service.Dashboard/web/templates/pages/settings"
)

type WebHandlers struct {
	AccountHandlers  *account.Handlers
	DomainHandlers   *domain.Handlers
	HomeHandlers     *home.Handlers
	HubHandlers      *hub.Handlers
	LinkHandlers     *link.Handlers
	SettingsHandlers *settings.Handlers
	Logger           *slog.Logger
	Authorizer       *extensions.Authorizer
	Ent              *ent.Client
}

func NewWebHandlers(logger *slog.Logger, authorizer *extensions.Authorizer, ent *ent.Client) *WebHandlers {

	return &WebHandlers{
		Logger:           logger,
		Authorizer:       authorizer,
		Ent:              ent,
		HubHandlers:      hub.NewHandlers(logger, ent),
		DomainHandlers:   domain.NewHandlers(logger, ent),
		AccountHandlers:  account.NewHandlers(),
		HomeHandlers:     home.NewHandlers(),
		LinkHandlers:     link.NewHandlers(logger, ent),
		SettingsHandlers: settings.NewHandlers(logger, ent),
	}
}
