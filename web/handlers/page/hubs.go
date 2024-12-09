package page

import (
	"log/slog"
	"net/http"
	"strings"

	"RouteHub.Service.Dashboard/ent"
	entHub "RouteHub.Service.Dashboard/ent/hub"
	"RouteHub.Service.Dashboard/ent/organization"
	"RouteHub.Service.Dashboard/ent/schema/enums"
	hub_enums "RouteHub.Service.Dashboard/ent/schema/enums/hub"
	"RouteHub.Service.Dashboard/ent/schema/mixin"
	"RouteHub.Service.Dashboard/features/hubConnection"
	"RouteHub.Service.Dashboard/web/extensions"
	"RouteHub.Service.Dashboard/web/templates/pages"
	"RouteHub.Service.Dashboard/web/templates/pages/components/domain"
	"RouteHub.Service.Dashboard/web/templates/pages/components/hub"
	"RouteHub.Service.Dashboard/web/templates/pages/components/utils"
	. "github.com/ahmetb/go-linq/v3"

	"github.com/labstack/echo/v4"
)

func (ph PageHandler) HubsHandler(c echo.Context) error {
	slog.InfoContext(c.Request().Context(), "Accessing index page", slog.String("method", "HubsHandler"))
	_, err := ph.GetUserInfo(c)

	if err != nil {
		return err
	}

	userInfo, userOrganization, err := ph.GetUserWithOrganization(c)

	if err != nil {
		return err
	}
	ph.Logger.Info("User Organization", "data", userOrganization)

	hubs, err := ph.Ent.Hub.Query().
		Where(entHub.HasOrganizationWith(organization.ID(userOrganization.ID))).
		WithDomain().
		All(c.Request().Context())

	if err != nil {
		ph.Logger.Error("Error fetching hubs", "error", err)
		hubs = []*ent.Hub{}
	}

	ph.Logger.Info("Hubs", "data", hubs)

	return extensions.Render(c, http.StatusOK, pages.Hubs(userInfo, hubs))
}

func (ph PageHandler) AttachHubGet(c echo.Context) error {
	isHXRequest := c.Request().Header.Get("HX-Request") != ""
	if !isHXRequest {
		return c.Redirect(http.StatusFound, "/hubs")
	}

	_, organization, err := ph.GetUserWithOrganization(c)
	if err != nil {
		msg := strings.Join([]string{"Error fetching user organization", err.Error()}, " ")
		feedback := utils.FormFeedback("error", nil, &msg)
		return extensions.Render(c, http.StatusOK, domain.CreateDomain(feedback, false))
	}

	domains, err := ph.Ent.Organization.QueryDomains(organization).All(c.Request().Context())
	if err != nil || len(domains) == 0 {
		msg := "Error fetching organization domains \nPlease add at least one domain"
		feedback := utils.FormFeedback("error", nil, &msg)
		return extensions.Render(c, http.StatusOK, domain.CreateDomain(feedback, false))
	}

	return extensions.Render(c, http.StatusOK, hub.AttachHub(domains, nil, true))
}

func (ph PageHandler) AttachHubPost(c echo.Context) error {
	title := "Attach Hub"

	address := c.FormValue("address")
	name := c.FormValue("name")
	slug := c.FormValue("slug")
	domain_id := c.FormValue("domain_id")

	_, organization, err := ph.GetUserWithOrganization(c)

	if err != nil {
		msg := strings.Join([]string{"Error fetching user organization", err.Error()}, " ")
		feedback := utils.FormFeedback("error", &title, &msg)
		return extensions.Render(c, http.StatusOK, hub.AttachHub(nil, feedback, true))
	}

	domains, err := ph.Ent.Organization.QueryDomains(organization).All(c.Request().Context())

	mqc, err := hubConnection.TryConnectToHub(address)
	if err != nil {
		msg := strings.Join([]string{"Error connecting to hub", err.Error()}, " ")
		feedback := utils.FormFeedback("error", &title, &msg)
		return extensions.Render(c, http.StatusOK, hub.AttachHub(domains, feedback, true))
	}

	ph.Logger.Info("MQTT Connection", "data", mqc.ClientID)

	domain := From(domains).WhereT(func(d *ent.Domain) bool {
		return d.ID == mixin.ID(domain_id)
	}).First().(*ent.Domain)

	if domain == nil {
		msg := "Domain not found"
		feedback := utils.FormFeedback("error", &title, &msg)
		return extensions.Render(c, http.StatusOK, hub.AttachHub(domains, feedback, true))
	}

	createdHub, err := ph.Ent.Hub.Create().
		SetName(name).
		SetSlug(slug).
		SetTCPAddress(address).
		SetOrganizationID(organization.ID).
		SetStatus(enums.StatusActive).
		SetDefaultRedirection(hub_enums.Timed).
		SetDomain(domain).
		Save(c.Request().Context())

	if err != nil {
		msg := strings.Join([]string{"Error creating hub", err.Error()}, " ")
		feedback := utils.FormFeedback("error", &title, &msg)
		return extensions.Render(c, http.StatusOK, hub.AttachHub(domains, feedback, true))
	}

	//c.Response().Header().Set("HX-Redirect", "/hubs")
	message := "Hub created successfully"
	feedback := utils.FormFeedback("success", &title, &message)

	ph.Logger.Info("Hub created", "hub", createdHub)
	return extensions.Render(c, http.StatusOK, hub.AttachHub(domains, feedback, false))

}

func (ph PageHandler) HubHandler(c echo.Context) error {
	user, err := ph.GetUserInfo(c)
	if err != nil {
		return c.Redirect(http.StatusForbidden, "/")
	}

	hub, err := ph.GetHubFromSlug(c, ph.Ent.Hub.Query().WithDomain().WithOrganization())

	if err != nil {
		ph.Logger.Error("Error fetching hub", "error", err)
		return c.Redirect(http.StatusFound, "/hubs")
	}

	ph.Logger.Info("Hub", "data", hub)

	return extensions.Render(c, http.StatusOK, pages.Hub(user, hub))
}
