package page

import (
	"net/http"

	"RouteHub.Service.Dashboard/ent/person"
	domainEnum "RouteHub.Service.Dashboard/ent/schema/enums/domain"
	"RouteHub.Service.Dashboard/web/extensions"
	"RouteHub.Service.Dashboard/web/templates/pages"
	"RouteHub.Service.Dashboard/web/templates/pages/components/domain"
	"RouteHub.Service.Dashboard/web/templates/pages/components/utils"
	"github.com/RouteHub-Link/DomainUtils/validator"
	"github.com/labstack/echo/v4"
)

func (ph PageHandler) DomainsHandler(c echo.Context) error {
	userInfo, _ := ph.GetUserInfo(c)

	personOrganizaitonAndDomains, err := ph.Ent.Person.Query().
		Where(person.SubjectID(userInfo.Subject)).
		WithOrganizations().
		QueryOrganizations().
		WithDomains().
		First(c.Request().Context())

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	domains := personOrganizaitonAndDomains.Edges.Domains

	ph.Logger.Info("Domains", "data", domains)

	return extensions.Render(c, http.StatusOK, pages.Domains(userInfo, domains))
}

func (ph PageHandler) CreateDomainGet(c echo.Context) error {
	isHXRequest := c.Request().Header.Get("HX-Request") != ""
	if !isHXRequest {
		return c.Redirect(http.StatusFound, "/domains")
	}

	return extensions.Render(c, http.StatusOK, domain.CreateDomain(nil, true))
}

func (ph PageHandler) CreateDomainPost(c echo.Context) error {
	title := "Create Domain"
	message := "Domain created successfully."

	url := c.FormValue("url")
	name := c.FormValue("name")
	ph.Logger.Info("URL", "url", url)

	var _validator = validator.DefaultValidator()
	isValid, err := _validator.ValidateURL(url)
	if !isValid || err != nil {
		title = "Invalid URL"
		message = "Please enter a valid URL; <br> " + err.Error()
		feedback := utils.FormFeedback("error", &title, &message)
		return extensions.Render(c, http.StatusOK, domain.CreateDomain(feedback, true))
	}

	userInfo, _ := ph.GetUserInfo(c)
	organizationsIds := ph.Ent.Person.Query().
		Where(person.SubjectID(userInfo.Subject)).
		WithOrganizations().
		QueryOrganizations().IDsX(c.Request().Context())

	if len(organizationsIds) == 0 {
		title = "No Organization"
		message = "You must be a member of an organization to create a domain."
		feedback := utils.FormFeedback("error", &title, &message)
		return extensions.Render(c, http.StatusOK, domain.CreateDomain(feedback, true))
	} else if len(organizationsIds) > 1 {
		title = "Multiple Organizations"
		message = "You are a member of multiple organizations. Please select one to create a domain."
		feedback := utils.FormFeedback("error", &title, &message)
		return extensions.Render(c, http.StatusOK, domain.CreateDomain(feedback, true))
	}

	// Create domain
	_, err = ph.Ent.Domain.Create().
		SetURL(url).
		SetName(name).
		SetOrganizationID(organizationsIds[0]).
		SetStatus(domainEnum.Passive).
		Save(c.Request().Context())

	if err != nil {
		title = "Error"
		message = "An error occurred while creating the domain; <br> " + err.Error()
		feedback := utils.FormFeedback("error", &title, &message)
		return extensions.Render(c, http.StatusOK, domain.CreateDomain(feedback, true))
	}

	c.Response().Header().Set("HX-Redirect", "/domains")
	return c.String(http.StatusOK, "Domain created successfully.")

}
