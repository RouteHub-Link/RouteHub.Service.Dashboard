package link

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"RouteHub.Service.Dashboard/ent"
	entLink "RouteHub.Service.Dashboard/ent/link"
	"RouteHub.Service.Dashboard/ent/schema/enums"
	"RouteHub.Service.Dashboard/ent/schema/types"
	"RouteHub.Service.Dashboard/features/scrape"
	"RouteHub.Service.Dashboard/web/context"
	"RouteHub.Service.Dashboard/web/extensions"
	"RouteHub.Service.Dashboard/web/templates/pages/hub/link/components"
	"RouteHub.Service.Dashboard/web/templates/pages/partial"
	"RouteHub.Service.Dashboard/web/utils"
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

var (
	collyClient = scrape.NewCollyClient()
)

func (h Handlers) HubLinksHandler(c echo.Context) error {
	user, err := context.GetUserFromContext(c)
	if err != nil {
		return c.Redirect(http.StatusForbidden, "/")
	}

	hub, err := context.GetHubFromContext(c)

	links, err := h.Ent.Hub.QueryLinks(hub).All(c.Request().Context())

	if err != nil {
		h.Logger.Error("Error fetching hub", "error", err)
		return c.Redirect(http.StatusFound, "/hubs")
	}

	return extensions.Render(c, http.StatusOK, index(user, hub, links))
}

func (h Handlers) HubLinksCreateHandler(c echo.Context) error {
	randSlug := utils.RandomShortKey()
	hubSlug := c.Param("slug")

	return extensions.Render(c, http.StatusOK, components.CreateLink(hubSlug, randSlug, nil, true))
}

type HubLinkCreate struct {
	Target string `json:"target_address" form:"target_address" validate:"required,url"`
	Slug   string `json:"slug" form:"slug" validate:"required"`
}

func (h Handlers) HubLinkCreatePostHandler(c echo.Context) error {
	title := "Short an URL"

	linkCreateRequest := new(HubLinkCreate)

	if err := extensions.BindAndValidate(c, linkCreateRequest); err != nil {
		msg := strings.Join([]string{"Error Validating Data", err.Error()}, " ")
		feedback := partial.FormFeedback("error", &title, &msg)

		randSlug := utils.RandomShortKey()
		hubSlug := c.Param("slug")

		return extensions.Render(c, http.StatusOK, components.CreateLink(hubSlug, randSlug, feedback, true))
	}

	hub, err := context.GetHubFromContext(c)
	if err != nil {
		h.Logger.Error("Error fetching hub", "error", err)
		c.Response().Header().Set("HX-Redirect", "/hubs")
		return err
	}

	scraped, err := collyClient.VisitScrapeOG(linkCreateRequest.Target)
	h.Logger.Info("Scraped", "scraped", scraped, "error", err)

	if err != nil {
		msg := strings.Join([]string{"Error Scraping URL", err.Error()}, " ")
		feedback := partial.FormFeedback("error", &title, &msg)
		return extensions.Render(c, http.StatusOK, components.CreateLink(hub.Slug, linkCreateRequest.Slug, feedback, true))
	}

	createQuery := h.Ent.Link.Create().
		SetPath(linkCreateRequest.Slug).
		SetTarget(linkCreateRequest.Target).
		SetHub(hub).
		SetRedirectionChoice(hub.DefaultRedirection).
		SetStatus(enums.StatusInactive)

	if scraped != nil {
		linkContent := types.LinkContent{
			MetaDescription: scraped,
			Title:           scraped.Title,
		}
		createQuery.SetLinkContent(linkContent)
	}

	link, err := createQuery.Save(c.Request().Context())

	if err != nil {
		h.Logger.Error("Error creating link", "error", err)
		msg := strings.Join([]string{"Error Creating Link", err.Error()}, " ")
		feedback := partial.FormFeedback("error", &title, &msg)
		return extensions.Render(c, http.StatusOK, components.CreateLink(hub.Slug, linkCreateRequest.Slug, feedback, true))
	}

	c.Response().Header().Set("HX-Redirect", fmt.Sprintf("/hubs/%s/links/%s", hub.Slug, link.Path))
	return nil
}

func (h Handlers) HubLinkEditHandler(c echo.Context) error {
	linkPath := c.Param("path")
	link, err := h.Ent.Link.Query().Where(entLink.PathEQ(linkPath)).Only(c.Request().Context())
	if err != nil {
		h.Logger.Error("Error fetching link", "error", err)
		return c.Redirect(http.StatusFound, "/links")
	}

	userInfo, _ := context.GetUserFromContext(c)
	hub, _ := context.GetHubFromContext(c)

	hubLinkPaylod := new(EditPayload)
	hubLinkPaylod.FromModel(link)

	return extensions.Render(c, http.StatusOK, edit(userInfo, hub, *hubLinkPaylod))
}
