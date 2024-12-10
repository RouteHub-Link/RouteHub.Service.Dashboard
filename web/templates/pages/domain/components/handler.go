package components

import (
	"net/http"

	domainEnum "RouteHub.Service.Dashboard/ent/schema/enums/domain"
	"RouteHub.Service.Dashboard/web/extensions"
	"RouteHub.Service.Dashboard/web/handlers/page"
	"RouteHub.Service.Dashboard/web/templates/pages/partial"
	"github.com/RouteHub-Link/DomainUtils/validator"
	"github.com/labstack/echo/v4"
)

type Handlers struct {
	RequestHandler *page.PageRequestHandler
}

func (h Handlers) DomainCreateGet(c echo.Context) error {
	isHXRequest := c.Request().Header.Get("HX-Request") != ""
	if !isHXRequest {
		return c.Redirect(http.StatusFound, "/domains")
	}

	return extensions.Render(c, http.StatusOK, CreateDomain(nil, true))
}

func (h Handlers) DomainCreatePost(c echo.Context) error {

	url := c.FormValue("url")
	name := c.FormValue("name")
	h.RequestHandler.Logger.Info("URL", "url", url)

	var _validator = validator.DefaultValidator()
	isValid, err := _validator.ValidateURL(url)
	if !isValid || err != nil {
		title := "Invalid URL"
		message := "Please enter a valid URL; <br> " + err.Error()
		feedback := partial.FormFeedback("error", &title, &message)
		return extensions.Render(c, http.StatusOK, CreateDomain(feedback, true))
	}

	_, organization, err := h.RequestHandler.GetUserWithOrganization(c)

	if err != nil {
		msg := err.Error()
		feedback := partial.FormFeedback("error", nil, &msg)
		return extensions.Render(c, http.StatusOK, CreateDomain(feedback, true))
	}

	// Create domain
	_, err = h.RequestHandler.Ent.Domain.Create().
		SetURL(url).
		SetName(name).
		SetOrganizationID(organization.ID).
		SetStatus(domainEnum.Passive).
		Save(c.Request().Context())

	if err != nil {
		title := "Error"
		message := "An error occurred while creating the domain; <br> " + err.Error()
		feedback := partial.FormFeedback("error", &title, &message)
		return extensions.Render(c, http.StatusOK, CreateDomain(feedback, true))
	}

	c.Response().Header().Set("HX-Redirect", "/domains")
	return c.String(http.StatusOK, "Domain created successfully.")

}
