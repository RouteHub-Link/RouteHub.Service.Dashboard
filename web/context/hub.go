package context

import (
	"net/http"

	"RouteHub.Service.Dashboard/ent"
	entHub "RouteHub.Service.Dashboard/ent/hub"
	"github.com/labstack/echo/v4"
)

const hubContextKey = "hub"

func GetHubFromContext(c echo.Context) (*ent.Hub, error) {
	hub, ok := c.Get(hubContextKey).(*ent.Hub)
	if !ok {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Hub not found in context")
	}

	return hub, nil
}

func HubContextMiddleware(ent *ent.Client) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			hub, err := getHubFromSlug(c, ent)
			if err != nil {
				c.Logger().Warn("getHubFromSlug throw an err", err)
			}

			c.Set(hubContextKey, hub)
			return next(c)
		}
	}
}

func getHubFromSlug(c echo.Context, ent *ent.Client) (*ent.Hub, error) {
	hubSlug := c.Param("slug")
	if hubSlug == "" || len(hubSlug) < 2 || len(hubSlug) > 100 {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid hub slug")
	}

	hub, err := ent.Hub.Query().WithDomain().
		Where(entHub.Slug(hubSlug)).Only(c.Request().Context())

	if err != nil {
		return nil, err
	}

	return hub, nil
}
