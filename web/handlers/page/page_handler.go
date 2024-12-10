package page

import (
	"log/slog"
	"net/http"

	"RouteHub.Service.Dashboard/ent"
	entHub "RouteHub.Service.Dashboard/ent/hub"
	"RouteHub.Service.Dashboard/ent/person"
	"RouteHub.Service.Dashboard/web/extensions"
	"github.com/labstack/echo/v4"
	"github.com/zitadel/oidc/v3/pkg/oidc"
)

type PageRequestHandler struct {
	Logger     *slog.Logger
	Authorizer *extensions.Authorizer
	Ent        *ent.Client
}

func NewPageHandler(logger *slog.Logger, authorizer *extensions.Authorizer, ent *ent.Client) *PageRequestHandler {
	return &PageRequestHandler{
		Logger:     logger,
		Authorizer: authorizer,
		Ent:        ent,
	}
}

func (h PageRequestHandler) GetUserInfo(c echo.Context) (*oidc.UserInfo, error) {

	var userInfo *oidc.UserInfo
	data, err := h.Authorizer.AUTHN.IsAuthenticated(c.Request())

	if err != nil {
		return nil, err
	}

	userInfo = data.GetUserInfo()

	h.Logger.Info("User info", "data", userInfo)

	return userInfo, nil
}

func (h PageRequestHandler) GetUserWithOrganization(c echo.Context) (*oidc.UserInfo, *ent.Organization, error) {
	userInfo, err := h.GetUserInfo(c)

	if err != nil {
		return nil, nil, err
	}

	organizationsIDs, err := h.Ent.Person.Query().
		Where(person.SubjectID(userInfo.Subject)).
		WithOrganizations().
		QueryOrganizations().
		IDs(c.Request().Context())

	if err != nil {
		return nil, nil, err
	} else if len(organizationsIDs) == 0 {
		return nil, nil, echo.NewHTTPError(http.StatusBadRequest, "You must be a member of an organization to create a domain.")
	} else if len(organizationsIDs) > 1 {
		return nil, nil, echo.NewHTTPError(http.StatusBadRequest, "You are a member of multiple organizations. Please select one to create a domain.")
	}

	organization, err := h.Ent.Organization.Get(c.Request().Context(), organizationsIDs[0])
	if err != nil {
		return nil, nil, err
	}

	return userInfo, organization, nil
}

func (h PageRequestHandler) GetHubFromSlug(c echo.Context, query *ent.HubQuery) (*ent.Hub, error) {
	hubSlug := c.Param("slug")
	if hubSlug == "" || len(hubSlug) < 2 || len(hubSlug) > 100 {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid hub slug")
	}

	if query == nil {
		query = h.Ent.Hub.Query()
	}

	query.Where(entHub.Slug(hubSlug))

	hub, err := query.WithDomain().Only(c.Request().Context())
	if err != nil {
		return nil, err
	}

	return hub, nil
}

func (h PageRequestHandler) BindAndValidate(c echo.Context, i interface{}) error {
	if err := c.Bind(i); err != nil {
		return err
	}

	if err := c.Validate(i); err != nil {
		return err
	}

	return nil
}
