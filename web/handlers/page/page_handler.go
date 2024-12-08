package page

import (
	"errors"
	"log/slog"

	"RouteHub.Service.Dashboard/ent"
	"RouteHub.Service.Dashboard/ent/person"
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

func (ph PageHandler) GetUserWithOrganization(c echo.Context) (*oidc.UserInfo, *ent.Organization, error) {
	userInfo, err := ph.GetUserInfo(c)

	if err != nil {
		return nil, nil, err
	}

	organizationsIDs, err := ph.Ent.Person.Query().
		Where(person.SubjectID(userInfo.Subject)).
		WithOrganizations().
		QueryOrganizations().IDs(c.Request().Context())

	if err != nil {
		return nil, nil, err
	} else if len(organizationsIDs) == 0 {
		message := "You must be a member of an organization to create a domain."
		return nil, nil, errors.New(message)
	} else if len(organizationsIDs) > 1 {
		message := "You are a member of multiple organizations. Please select one to create a domain."
		return nil, nil, errors.New(message)
	}

	organization, err := ph.Ent.Organization.Get(c.Request().Context(), organizationsIDs[0])
	if err != nil {
		return nil, nil, err
	}

	return userInfo, organization, nil
}
