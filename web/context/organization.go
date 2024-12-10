package context

import (
	"net/http"

	"RouteHub.Service.Dashboard/ent"
	"RouteHub.Service.Dashboard/ent/person"
	"github.com/labstack/echo/v4"
)

const OrganizationContextKey = "organization"

func GetOrganizationFromContext(c echo.Context) (*ent.Organization, error) {
	organization, ok := c.Get(OrganizationContextKey).(*ent.Organization)
	if !ok {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "organization not found in context")
	}

	return organization, nil
}

func OrganizationContextMiddleware(ent *ent.Client) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user, err := GetUserFromContext(c)
			if err != nil {
				c.Logger().Warn("getUserFromContext throw an err", err)
				return err
			}

			organization, err := getOrganizationFromContext(c, ent, user.Subject)
			if err != nil {
				c.Logger().Warn("getOrganizationFromUser throw an err", err)
			}

			c.Set(OrganizationContextKey, organization)
			return next(c)
		}
	}
}

func getOrganizationFromContext(c echo.Context, ent *ent.Client, subjectID string) (org *ent.Organization, err error) {
	if org, err = ent.Person.
		Query().
		Where(person.SubjectID(subjectID)).
		WithOrganizations().
		QueryOrganizations().
		First(c.Request().Context()); err != nil {
		return nil, err
	}

	return
}
