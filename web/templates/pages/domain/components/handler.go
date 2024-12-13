package components

import (
	"log/slog"
	"net/http"

	"RouteHub.Service.Dashboard/ent"
	"RouteHub.Service.Dashboard/ent/person"
	domainEnum "RouteHub.Service.Dashboard/ent/schema/enums/domain"
	"RouteHub.Service.Dashboard/web/context"
	"RouteHub.Service.Dashboard/web/extensions"
	"RouteHub.Service.Dashboard/web/templates/pages/partial"
	"github.com/RouteHub-Link/DomainUtils/validator"
	"github.com/labstack/echo/v4"
)

type Handlers struct {
	Logger *slog.Logger
	Ent    *ent.Client
}

func NewHandlers(logger *slog.Logger, ent *ent.Client) *Handlers {
	return &Handlers{
		Logger: logger,
		Ent:    ent,
	}
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
	h.Logger.Info("URL", "url", url)

	var _validator = validator.DefaultValidator()
	isValid, err := _validator.ValidateURL(url)
	if !isValid || err != nil {
		title := "Invalid URL"
		message := "Please enter a valid URL; <br> " + err.Error()
		feedback := partial.FormFeedback("error", &title, &message)
		return extensions.Render(c, http.StatusOK, CreateDomain(feedback, true))
	}

	organization, err := context.GetOrganizationFromContext(c)

	if err != nil {
		msg := err.Error()
		feedback := partial.FormFeedback("error", nil, &msg)
		return extensions.Render(c, http.StatusOK, CreateDomain(feedback, true))
	}

	// Create domain
	_, err = h.Ent.Domain.Create().
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

	message := "Domain created successfully."
	feedback := partial.FormFeedback("success", nil, &message)

	extensions.HTMXAppendTrigger(c, "newDomain, close-modal")
	extensions.HTMXAppendSuccessToast(c, message)

	return extensions.Render(c, http.StatusOK, CreateDomain(feedback, false))
}

func (h Handlers) PartialDomainTable(c echo.Context) error {
	userInfo, _ := context.GetUserFromContext(c)

	personOrganizaitonAndDomains, err := h.Ent.Person.Query().
		Where(person.SubjectID(userInfo.Subject)).
		WithOrganizations().
		QueryOrganizations().
		WithDomains().
		First(c.Request().Context())

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	domains := personOrganizaitonAndDomains.Edges.Domains

	return extensions.Render(c, http.StatusOK, domainsTableContainer(domains))
}
