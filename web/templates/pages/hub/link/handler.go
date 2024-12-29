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

	hub, _ := context.GetHubFromContext(c)

	links, err := h.Ent.Hub.QueryLinks(hub).All(c.Request().Context())

	if err != nil {
		h.Logger.Error("Error fetching hub", "error", err)
		return c.Redirect(http.StatusFound, "/hubs")
	}

	return extensions.Render(c, http.StatusOK, index(user, hub, links))
}

func (h Handlers) HubLinksCreate(c echo.Context) error {
	randSlug := utils.RandomShortKey()
	hubSlug := c.Param("slug")

	return extensions.Render(c, http.StatusOK, components.CreateLink(hubSlug, randSlug, nil, true))
}

type HubLinkCreate struct {
	Target string `json:"target_address" form:"target_address" validate:"required,url"`
	Slug   string `json:"slug" form:"slug" validate:"required"`
}

func (h Handlers) HubLinkCreatePost(c echo.Context) error {
	title := "Short an URL"

	linkCreateRequest := new(HubLinkCreate)

	if err := extensions.BindAndValidate(c, linkCreateRequest); err != nil {
		msg := strings.Join([]string{"Error Validating Data", err.Error()}, " ")
		feedback := partial.FormFeedback("error", title, msg)

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
		feedback := partial.FormFeedback("error", title, msg)
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
		createQuery.SetLinkContent(&linkContent)
	}

	link, err := createQuery.Save(c.Request().Context())

	if err != nil {
		h.Logger.Error("Error creating link", "error", err)
		msg := strings.Join([]string{"Error Creating Link", err.Error()}, " ")
		feedback := partial.FormFeedback("error", title, msg)
		return extensions.Render(c, http.StatusOK, components.CreateLink(hub.Slug, linkCreateRequest.Slug, feedback, true))
	}

	c.Response().Header().Set("HX-Redirect", fmt.Sprintf("/hub/%s/links/%s", hub.Slug, link.Path))
	return nil
}

func (h Handlers) HubLinkEditGet(c echo.Context) error {
	linkPath := c.Param("path")
	link, err := h.Ent.Link.Query().Where(entLink.PathEQ(linkPath)).Only(c.Request().Context())
	if err != nil {
		h.Logger.Error("Error fetching link", "error", err)
		return c.Redirect(http.StatusFound, "/links")
	}

	userInfo, _ := context.GetUserFromContext(c)
	hub, _ := context.GetHubFromContext(c)

	hubLinkPaylod := new(EditLinkPayload)
	hubLinkPaylod.FromModel(link)

	hubLinkMetaPayload := new(EditLinkMetaDescriptionPayload)
	hubLinkMetaPayload.FromModel(link.LinkContent.MetaDescription)

	return extensions.Render(c, http.StatusOK, edit(userInfo, hub, link, *hubLinkPaylod, *hubLinkMetaPayload))
}

func (h Handlers) HubLinkEditPost(c echo.Context) error {
	title := "Edit Link"
	linkPath := c.Param("path")
	link, err := h.Ent.Link.Query().Where(entLink.PathEQ(linkPath)).Only(c.Request().Context())

	if err != nil {
		h.Logger.Error("Error fetching link", "error", err)
		return c.Redirect(http.StatusFound, "/links")
	}

	LinkUpdatePayload := new(EditLinkPayload)
	MetaPayload := new(EditLinkMetaDescriptionPayload)

	hub, _ := context.GetHubFromContext(c)

	if err := extensions.BindAndValidate(c, LinkUpdatePayload); err != nil {
		msg := strings.Join([]string{"Error Validating Data", err.Error()}, " ")
		feedback := partial.FormFeedback("error", title, msg)
		return extensions.Render(c, http.StatusOK, editForm(hub, link, *LinkUpdatePayload, *MetaPayload, feedback))
	}

	if err := extensions.BindAndValidate(c, MetaPayload); err != nil {
		msg := strings.Join([]string{"Error Validating Data", err.Error()}, " ")
		feedback := partial.FormFeedback("error", title, msg)
		return extensions.Render(c, http.StatusOK, editForm(hub, link, *LinkUpdatePayload, *MetaPayload, feedback))
	}

	isPathChanged := linkPath != LinkUpdatePayload.Path

	bucketPath := strings.Join([]string{"hubs", hub.Slug, "links", link.Path}, "/")
	err = handleFavIcon(MetaPayload, link, c, bucketPath)
	if err != nil {
		feedback := partial.FormFeedbackFromErr(title, err)
		return extensions.Render(c, http.StatusOK, editForm(hub, link, *LinkUpdatePayload, *MetaPayload, feedback))
	}

	err = handleOGImage(MetaPayload, link, c, bucketPath)
	if err != nil {
		feedback := partial.FormFeedbackFromErr(title, err)
		return extensions.Render(c, http.StatusOK, editForm(hub, link, *LinkUpdatePayload, *MetaPayload, feedback))
	}

	LinkUpdatePayload.UpdateModel(link)
	if link.LinkContent.MetaDescription == nil {
		link.LinkContent.MetaDescription = &types.MetaDescription{}
	}

	MetaPayload.UpdateModel(link.LinkContent.MetaDescription)

	link, err = h.Ent.Link.UpdateOne(link).
		SetLinkContent(link.LinkContent).
		SetPath(LinkUpdatePayload.Path).
		SetTarget(LinkUpdatePayload.Target).
		SetRedirectionChoice(LinkUpdatePayload.RedirectionChoice).
		//SetStatus(LinkUpdatePayload.Status).
		Save(c.Request().Context())

	if err != nil {
		h.Logger.Error("Error updating link", "error", err)
		msg := strings.Join([]string{"Error Updating Link", err.Error()}, " ")
		feedback := partial.FormFeedback("error", title, msg)
		return extensions.Render(c, http.StatusOK, editForm(hub, link, *LinkUpdatePayload, *MetaPayload, feedback))
	}

	if isPathChanged {
		// HX-Replace-Url
		c.Response().Header().Set("HX-Replace-Url", fmt.Sprintf("/hubs/%s/links/%s", hub.Slug, link.Path))
	}

	extensions.HTMXAppendSuccessToast(c, "Link Updated Successfully")
	extensions.HTMXAppendPrelineRefresh(c)

	LinkUpdatePayload.FromModel(link)
	MetaPayload.FromModel(link.LinkContent.MetaDescription)

	return extensions.Render(c, http.StatusOK, editForm(hub, link, *LinkUpdatePayload, *MetaPayload, nil))
}

func handleFavIcon(MetaPayload *EditLinkMetaDescriptionPayload, link *ent.Link, c echo.Context, bucketPath string) error {
	if err := extensions.ProcessFileFromEchoContext(c, &MetaPayload.FavIcon, "meta_description_favicon", bucketPath, "favicon"); err != nil {
		msg := strings.Join([]string{"Error Processing File", err.Error()}, " ")
		return fmt.Errorf("%s", msg)
	} else {
		MetaPayload.FavIcon = utils.LinkToS3Path(MetaPayload.FavIcon)
	}
	return nil
}

func handleOGImage(MetaPayload *EditLinkMetaDescriptionPayload, link *ent.Link, c echo.Context, bucketPath string) error {
	if err := extensions.ProcessFileFromEchoContext(c, &MetaPayload.OGBigImage, "meta_description_og_big_image", bucketPath, "og_image"); err != nil {
		msg := strings.Join([]string{"Error Processing File", err.Error()}, " ")
		return fmt.Errorf("%s", msg)
	} else {
		MetaPayload.OGBigImage = utils.LinkToS3Path(MetaPayload.OGBigImage)
	}
	return nil
}

func (h Handlers) HubLinkStatusGet(c echo.Context) error {
	linkPath := c.Param("path")
	link, err := h.Ent.Link.Query().Where(entLink.PathEQ(linkPath)).Only(c.Request().Context())
	if err != nil {
		h.Logger.Error("Error fetching link", "error", err)
		return c.Redirect(http.StatusFound, "/links")
	}

	hub, _ := context.GetHubFromContext(c)

	payload := components.EditLinkStatusPayload{
		Status:   link.Status,
		HubSlug:  hub.Slug,
		LinkPath: link.Path,
	}

	return extensions.Render(c, http.StatusOK, components.LinkStatusForm(payload))
}

func (h Handlers) HubLinkStatusPost(c echo.Context) error {
	linkPath := c.Param("path")
	link, err := h.Ent.Link.Query().Where(entLink.PathEQ(linkPath)).Only(c.Request().Context())
	if err != nil {
		h.Logger.Error("Error fetching link", "error", err)
		return c.Redirect(http.StatusFound, "/links")
	}

	payload := new(components.EditLinkStatusPayload)

	if err := extensions.BindAndValidate(c, payload); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	hub, _ := context.GetHubFromContext(c)
	payload.HubSlug = hub.Slug
	payload.LinkPath = link.Path

	link, err = h.Ent.Link.UpdateOne(link).
		SetStatus(payload.Status).
		Save(c.Request().Context())

	if err != nil {
		h.Logger.Error("Error updating link", "error", err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	extensions.HTMXAppendSuccessToast(c, "Link Status Updated Successfully")
	extensions.HTMXAppendPrelineRefresh(c)

	return extensions.Render(c, http.StatusOK, components.LinkStatusForm(*payload))
}
