package page

import (
	"log/slog"

	"RouteHub.Service.Dashboard/ent"
	"RouteHub.Service.Dashboard/web/extensions"
)

type PageHandler struct {
	Logger     *slog.Logger
	Authorizer *extensions.Authorizer
	Ent        *ent.Client
}

func NewPageHandler(logger *slog.Logger, authorizer *extensions.Authorizer, ent *ent.Client) *PageHandler {
	return &PageHandler{
		Logger:     logger,
		Authorizer: authorizer,
		Ent:        ent,
	}
}
