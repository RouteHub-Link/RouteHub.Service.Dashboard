package components

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"RouteHub.Service.Dashboard/ent"
	"RouteHub.Service.Dashboard/ent/schema/enums"
	"RouteHub.Service.Dashboard/ent/schema/mixin"
	"RouteHub.Service.Dashboard/features/hubConnection"
	"RouteHub.Service.Dashboard/web/context"
	"RouteHub.Service.Dashboard/web/extensions"
	domainComponents "RouteHub.Service.Dashboard/web/templates/pages/domain/components"
	"RouteHub.Service.Dashboard/web/templates/pages/partial"
	. "github.com/ahmetb/go-linq/v3"
	"github.com/charmbracelet/hotdiva2000"
	"github.com/labstack/echo/v4"
)

type Handlers struct {
	Ent    *ent.Client
	Logger *slog.Logger
}

func NewHandlers(logger *slog.Logger, ent *ent.Client) *Handlers {
	return &Handlers{
		Logger: logger,
		Ent:    ent,
	}
}

func (h Handlers) AttachHubGet(c echo.Context) error {
	isHXRequest := c.Request().Header.Get("HX-Request") != ""
	if !isHXRequest {
		return c.Redirect(http.StatusFound, "/hubs")
	}

	organization, err := context.GetOrganizationFromContext(c)
	if err != nil {
		msg := strings.Join([]string{"Error fetching user organization", err.Error()}, " ")
		feedback := partial.FormFeedback("error", nil, &msg)
		return extensions.Render(c, http.StatusOK, domainComponents.CreateDomain(feedback, false))
	}

	humanizedSlug := hotdiva2000.Generate()

	domains, err := h.Ent.Organization.QueryDomains(organization).All(c.Request().Context())
	if err != nil || len(domains) == 0 {
		msg := "Error fetching organization domains \nPlease add at least one domain"
		feedback := partial.FormFeedback("error", nil, &msg)
		return extensions.Render(c, http.StatusOK, domainComponents.CreateDomain(feedback, false))
	}

	return extensions.Render(c, http.StatusOK, AttachHub(humanizedSlug, domains, nil, true))
}

func (h Handlers) AttachHubPost(c echo.Context) error {
	title := "Attach Hub"

	address := c.FormValue("address")
	name := c.FormValue("name")
	slug := c.FormValue("slug")
	domainID := c.FormValue("domain_id")

	organization, err := context.GetOrganizationFromContext(c)

	if err != nil {
		msg := strings.Join([]string{"Error fetching user organization", err.Error()}, " ")
		feedback := partial.FormFeedback("error", &title, &msg)
		return extensions.Render(c, http.StatusOK, AttachHub(slug, nil, feedback, true))
	}

	domains, err := h.Ent.Organization.QueryDomains(organization).All(c.Request().Context())

	mqc, err := hubConnection.TryConnectToHub(address)
	if err != nil {
		msg := strings.Join([]string{"Error connecting to hub", err.Error()}, " ")
		feedback := partial.FormFeedback("error", &title, &msg)
		return extensions.Render(c, http.StatusOK, AttachHub(slug, domains, feedback, true))
	}

	h.Logger.Info("MQTT Connection", "data", mqc.ClientID)

	domain := From(domains).WhereT(func(d *ent.Domain) bool {
		return d.ID == mixin.ID(domainID)
	}).First().(*ent.Domain)

	if domain == nil {
		msg := "Domain not found"
		feedback := partial.FormFeedback("error", &title, &msg)
		return extensions.Render(c, http.StatusOK, AttachHub(slug, domains, feedback, true))
	}

	createdHub, err := h.Ent.Hub.Create().
		SetName(name).
		SetSlug(slug).
		SetTCPAddress(address).
		SetOrganizationID(organization.ID).
		SetStatus(enums.StatusActive).
		SetDefaultRedirection(enums.RedirectionChoiceTimed).
		SetDomain(domain).
		Save(c.Request().Context())

	if err != nil {
		msg := strings.Join([]string{"Error creating hub", err.Error()}, " ")
		feedback := partial.FormFeedback("error", &title, &msg)
		return extensions.Render(c, http.StatusOK, AttachHub(slug, domains, feedback, true))
	}

	c.Response().Header().Set("HX-Redirect", fmt.Sprintf("/hub/%s", createdHub.Slug))
	return extensions.Render(c, http.StatusOK, AttachHub(slug, domains, nil, false))
}
