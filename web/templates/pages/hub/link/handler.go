package link

import (
	"fmt"
	"net/http"
	"strings"

	"RouteHub.Service.Dashboard/ent/schema/enums"
	"RouteHub.Service.Dashboard/ent/schema/types"
	"RouteHub.Service.Dashboard/features/scrape"
	"RouteHub.Service.Dashboard/web/extensions"
	"RouteHub.Service.Dashboard/web/handlers/page"
	"RouteHub.Service.Dashboard/web/templates/pages/hub/link/components"
	"RouteHub.Service.Dashboard/web/templates/pages/partial"
	"RouteHub.Service.Dashboard/web/utils"
	"github.com/labstack/echo/v4"
)

type Handlers struct {
	RequestHandler *page.PageRequestHandler
}

func NewHandlers(RequestHandler *page.PageRequestHandler) *Handlers {
	return &Handlers{
		RequestHandler: RequestHandler,
	}
}

var (
	collyClient = scrape.NewCollyClient()
)

func (h Handlers) HubLinksHandler(c echo.Context) error {
	user, err := h.RequestHandler.GetUserInfo(c)
	if err != nil {
		return c.Redirect(http.StatusForbidden, "/")
	}

	query := h.RequestHandler.Ent.Hub.Query().WithDomain().WithLinks()

	hub, err := h.RequestHandler.GetHubFromSlug(c, query)

	if err != nil {
		h.RequestHandler.Logger.Error("Error fetching hub", "error", err)
		return c.Redirect(http.StatusFound, "/hubs")
	}

	return extensions.Render(c, http.StatusOK, index(user, hub, hub.Edges.Links))
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

	if err := h.RequestHandler.BindAndValidate(c, linkCreateRequest); err != nil {
		msg := strings.Join([]string{"Error Validating Data", err.Error()}, " ")
		feedback := partial.FormFeedback("error", &title, &msg)

		randSlug := utils.RandomShortKey()
		hubSlug := c.Param("slug")

		return extensions.Render(c, http.StatusOK, components.CreateLink(hubSlug, randSlug, feedback, true))
	}

	hub, err := h.RequestHandler.GetHubFromSlug(c, nil)
	if err != nil {
		h.RequestHandler.Logger.Error("Error fetching hub", "error", err)
		c.Response().Header().Set("HX-Redirect", "/hubs")
		return err
	}

	scraped, err := collyClient.VisitScrapeOG(linkCreateRequest.Target)
	h.RequestHandler.Logger.Info("Scraped", "scraped", scraped, "error", err)

	if err != nil {
		msg := strings.Join([]string{"Error Scraping URL", err.Error()}, " ")
		feedback := partial.FormFeedback("error", &title, &msg)
		return extensions.Render(c, http.StatusOK, components.CreateLink(hub.Slug, linkCreateRequest.Slug, feedback, true))
	}

	createQuery := h.RequestHandler.Ent.Link.Create().
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
		h.RequestHandler.Logger.Error("Error creating link", "error", err)
		msg := strings.Join([]string{"Error Creating Link", err.Error()}, " ")
		feedback := partial.FormFeedback("error", &title, &msg)
		return extensions.Render(c, http.StatusOK, components.CreateLink(hub.Slug, linkCreateRequest.Slug, feedback, true))
	}

	c.Response().Header().Set("HX-Redirect", fmt.Sprintf("/hubs/%s/links/%s", hub.Slug, link.Path))
	return nil
}
