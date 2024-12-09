package page

import (
	"fmt"
	"net/http"
	"strings"

	"RouteHub.Service.Dashboard/ent/schema/enums"
	"RouteHub.Service.Dashboard/ent/schema/types"
	"RouteHub.Service.Dashboard/features/scrape"
	"RouteHub.Service.Dashboard/web/extensions"
	"RouteHub.Service.Dashboard/web/templates/pages"
	linkComponents "RouteHub.Service.Dashboard/web/templates/pages/components/link"
	"RouteHub.Service.Dashboard/web/templates/pages/components/utils"
	"github.com/labstack/echo/v4"
)

var (
	collyClient = scrape.NewCollyClient()
)

func (ph PageHandler) HubLinksHandler(c echo.Context) error {
	user, err := ph.GetUserInfo(c)
	if err != nil {
		return c.Redirect(http.StatusForbidden, "/")
	}

	query := ph.Ent.Hub.Query().WithDomain().WithLinks()

	hub, err := ph.GetHubFromSlug(c, query)

	if err != nil {
		ph.Logger.Error("Error fetching hub", "error", err)
		return c.Redirect(http.StatusFound, "/hubs")
	}

	return extensions.Render(c, http.StatusOK, pages.HubLinks(user, hub, hub.Edges.Links))
}

func (ph PageHandler) HubLinksCreateHandler(c echo.Context) error {
	randSlug := utils.RandomShortKey()
	hubSlug := c.Param("slug")

	return extensions.Render(c, http.StatusOK, linkComponents.CreateLink(hubSlug, randSlug, nil, true))
}

type HubLinkCreate struct {
	Target string `json:"target_address" form:"target_address" validate:"required,url"`
	Slug   string `json:"slug" form:"slug" validate:"required"`
}

func (ph PageHandler) HubLinkCreatePostHandler(c echo.Context) error {
	title := "Short an URL"

	linkCreateRequest := new(HubLinkCreate)

	if err := ph.BindAndValidate(c, linkCreateRequest); err != nil {
		msg := strings.Join([]string{"Error Validating Data", err.Error()}, " ")
		feedback := utils.FormFeedback("error", &title, &msg)

		randSlug := utils.RandomShortKey()
		hubSlug := c.Param("slug")

		return extensions.Render(c, http.StatusOK, linkComponents.CreateLink(hubSlug, randSlug, feedback, true))
	}

	hub, err := ph.GetHubFromSlug(c, nil)
	if err != nil {
		ph.Logger.Error("Error fetching hub", "error", err)
		c.Response().Header().Set("HX-Redirect", "/hubs")
		return err
	}

	scraped, err := collyClient.VisitScrapeOG(linkCreateRequest.Target)
	ph.Logger.Info("Scraped", "scraped", scraped, "error", err)

	if err != nil {
		msg := strings.Join([]string{"Error Scraping URL", err.Error()}, " ")
		feedback := utils.FormFeedback("error", &title, &msg)
		return extensions.Render(c, http.StatusOK, linkComponents.CreateLink(hub.Slug, linkCreateRequest.Slug, feedback, true))
	}

	createQuery := ph.Ent.Link.Create().
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
		ph.Logger.Error("Error creating link", "error", err)
		msg := strings.Join([]string{"Error Creating Link", err.Error()}, " ")
		feedback := utils.FormFeedback("error", &title, &msg)
		return extensions.Render(c, http.StatusOK, linkComponents.CreateLink(hub.Slug, linkCreateRequest.Slug, feedback, true))
	}

	c.Response().Header().Set("HX-Redirect", fmt.Sprintf("/hubs/%s/links/%s", hub.Slug, link.Path))
	return nil
}
