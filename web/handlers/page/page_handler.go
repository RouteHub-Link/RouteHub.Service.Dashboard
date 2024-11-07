package page

import (
	"log/slog"

	"RouteHub.Service.Dashboard/ent"
	"RouteHub.Service.Dashboard/web/extensions"
	"github.com/labstack/echo/v4"
	"github.com/zitadel/oidc/v3/pkg/oidc"
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

func (ph PageHandler) GetUserInfo(c echo.Context) (*oidc.UserInfo, error) {

	var userInfo *oidc.UserInfo
	data, err := ph.Authorizer.AUTHN.IsAuthenticated(c.Request())

	if err != nil {
		return nil, err
	}

	userInfo = data.GetUserInfo()

	ph.Logger.Info("User info", "data", userInfo)

	return userInfo, nil
}
