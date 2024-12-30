package pages

import (
	"fmt"
	"log/slog"
	"net/http"

	"RouteHub.Service.Dashboard/ent"
	"RouteHub.Service.Dashboard/ent/hub"
	"RouteHub.Service.Dashboard/ent/page"
	entPage "RouteHub.Service.Dashboard/ent/page"
	"RouteHub.Service.Dashboard/ent/schema/enums"
	"RouteHub.Service.Dashboard/ent/schema/types"
	"RouteHub.Service.Dashboard/web/context"
	"RouteHub.Service.Dashboard/web/extensions"
	"RouteHub.Service.Dashboard/web/templates/pages/partial"
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

type PagePayload struct {
	PageSlug    string `json:"page_slug" form:"page_slug" validate:"required"`
	Name        string `json:"page_name" form:"page_name" validate:"required"`
	Description string `json:"page_description" form:"page_description" validate:"required"`
	ContentJSON string `json:"page_content_json" form:"page_content_json"`
	ContentHTML string `json:"page_content_html" form:"page_content_html"`
	MetaData    *types.MetaDescription
}

func (p *PagePayload) UpdateFromEntity(page *ent.Page) {
	p.PageSlug = page.Slug
	p.Name = page.Name
	p.Description = page.PageDescription
	p.ContentJSON = page.PageContentJSON
	p.ContentHTML = page.PageContentHTML
	p.MetaData = page.MetaDescription
}

const (
	EmptyPageData = `{"assets":[],"styles":[{"selectors":["#iqsl"],"style":{"background-color":"#000000"},"wrapper":1},{"selectors":[{"name":"gjs-row","private":1}],"style":{"display":"table","padding":"10px","width":"100%"}},{"selectors":[{"name":"gjs-cell","private":1}],"style":{"width":"100%","display":"block"},"mediaText":"(max-width: 768px)","atRuleType":"media"},{"selectors":["gjs-cell30"],"style":{"width":"100%","display":"block"},"mediaText":"(max-width: 768px)","atRuleType":"media"},{"selectors":["gjs-cell70"],"style":{"width":"100%","display":"block"},"mediaText":"(max-width: 768px)","atRuleType":"media"},{"selectors":[{"name":"gjs-cell","private":1}],"style":{"width":"8%","display":"table-cell","height":"75px"}},{"selectors":["#iide"],"style":{"height":"480px"}}],"pages":[{"frames":[{"component":{"type":"wrapper","stylable":["background","background-color","background-image","background-repeat","background-attachment","background-position","background-size"],"attributes":{"id":"iqsl"},"components":[{"name":"Row","droppable":".gjs-cell","resizable":{"tl":0,"tc":0,"tr":0,"cl":0,"cr":0,"bl":0,"br":0,"minDim":1},"classes":[{"name":"gjs-row","private":1}],"attributes":{"id":"iide"},"components":[{"name":"Cell","draggable":".gjs-row","resizable":{"tl":0,"tc":0,"tr":0,"cl":0,"cr":1,"bl":0,"br":0,"minDim":1,"bc":0,"currentUnit":1,"step":0.2},"classes":[{"name":"gjs-cell","private":1}],"attributes":{"id":"i1tr"},"components":[{"classes":["box"],"components":[{"type":"text","classes":["content"],"components":[{"type":"textnode","content":"Welcome the page editor. Please check out component's and layer editor on the right."}]}]}]}]}],"head":{"type":"head"},"docEl":{"tagName":"html"}},"id":"S2ky80NM1QN5Lb4V"}],"id":"mb8rWiK0aH25mHE0"}],"symbols":[],"dataSources":[]}`
)

func (h Handlers) IndexGetHandler(c echo.Context) error {
	userInfo, err := context.GetUserFromContext(c)
	if err != nil {
		return c.Redirect(http.StatusForbidden, "/")
	}

	contextHub, _ := context.GetHubFromContext(c)

	pages, err := h.Ent.Page.Query().
		Where(entPage.HasHubWith(hub.ID(contextHub.ID))).
		All(c.Request().Context())

	pagePayloads := make([]PagePayload, 0, len(pages))
	for _, page := range pages {
		payload := PagePayload{}
		payload.UpdateFromEntity(page)
		pagePayloads = append(pagePayloads, payload)
	}

	return extensions.Render(c, http.StatusOK, index(userInfo, contextHub, &pagePayloads))

}

func (h Handlers) NewGetHandler(c echo.Context) error {
	_, err := context.GetUserFromContext(c)
	if err != nil {
		return c.Redirect(http.StatusForbidden, "/")
	}

	hub, _ := context.GetHubFromContext(c)
	payload := CreatePagePayload{}

	return extensions.Render(c, http.StatusOK, createPage(hub.Slug, payload, nil))
}

func (h Handlers) NewPostHandler(c echo.Context) error {
	title := "Create Page"
	_, err := context.GetUserFromContext(c)
	if err != nil {
		return c.Redirect(http.StatusForbidden, "/")
	}

	hub, _ := context.GetHubFromContext(c)
	payload := CreatePagePayload{}

	if err := extensions.BindAndValidate(c, &payload); err != nil {
		feedback := partial.FormFeedbackFromErr(title, err)
		return extensions.Render(c, http.StatusOK, createPage(hub.Slug, payload, feedback))
	}

	if payload.Slug[0] == '/' {
		payload.Slug = payload.Slug[1:]
	}

	page, err := h.Ent.Page.Create().
		SetName(payload.Title).
		SetPageDescription(payload.Description).
		SetSlug(payload.Slug).
		SetMetaDescription(&types.MetaDescription{Title: payload.Title, Description: payload.Description}).
		SetHub(hub).
		SetStatus(enums.StatusInactive).
		Save(c.Request().Context())

	if err != nil {
		feedback := partial.FormFeedbackFromErr(title, err)
		return extensions.Render(c, http.StatusOK, createPage(hub.Slug, payload, feedback))
	}

	c.Response().Header().Set("HX-Redirect", fmt.Sprintf("/hub/%s/pages/%s", hub.Slug, page.Slug))
	extensions.HTMXAppendSuccessToast(c, "Page created successfully")
	extensions.HTMXAppendPrelineRefresh(c)

	return extensions.Render(c, http.StatusOK, createPage(hub.Slug, payload, nil))
}

func (h Handlers) EditGetHandler(c echo.Context) error {
	ctx := c.Request().Context()

	userInfo, err := context.GetUserFromContext(c)
	if err != nil {
		return c.Redirect(http.StatusForbidden, "/")
	}

	contextHub, _ := context.GetHubFromContext(c)
	if err != nil {
		return c.Redirect(http.StatusNotFound, "/hubs")
	}

	pageSlug := c.Param("pageSlug")

	page, err := h.Ent.Page.Query().
		Where(entPage.HasHubWith(hub.Slug(contextHub.Slug))).
		Where(page.Slug(pageSlug)).
		Only(ctx)

	if err != nil {
		return c.Redirect(http.StatusNotFound, "/hubs")
	}

	if page.PageContentJSON == "" {
		page.PageContentJSON = EmptyPageData
	}

	payload := new(PagePayload)
	payload.UpdateFromEntity(page)

	return extensions.Render(c, http.StatusOK, editPage(userInfo, contextHub, *payload))
}

func (h Handlers) EditFormGetHandler(c echo.Context) error {
	ctx := c.Request().Context()

	_, err := context.GetUserFromContext(c)
	if err != nil {
		feedback := partial.FormFeedbackFromErr("Edit Page", err)
		return extensions.Render(c, http.StatusOK, editForm("", PagePayload{}, feedback))
	}

	contextHub, err := context.GetHubFromContext(c)
	if err != nil {
		feedback := partial.FormFeedbackFromErr("Edit Page", err)
		return extensions.Render(c, http.StatusOK, editForm("", PagePayload{}, feedback))
	}

	pageSlug := c.Param("pageSlug")

	page, err := h.Ent.Page.Query().
		Where(entPage.HasHubWith(hub.ID(contextHub.ID))).
		Where(page.Slug(pageSlug)).
		Only(ctx)

	if err != nil {
		feedback := partial.FormFeedbackFromErr("Edit Page", err)
		return extensions.Render(c, http.StatusOK, editForm(contextHub.Slug, PagePayload{}, feedback))
	}

	if page.PageContentJSON == "" {
		page.PageContentJSON = EmptyPageData
	}

	payload := new(PagePayload)
	payload.UpdateFromEntity(page)

	return extensions.Render(c, http.StatusOK, editForm(contextHub.Slug, *payload, nil))

}

func (h Handlers) EditPostHandler(c echo.Context) error {
	ctx := c.Request().Context()

	_, err := context.GetUserFromContext(c)
	if err != nil {
		return c.Redirect(http.StatusForbidden, "/")
	}

	contextHub, err := context.GetHubFromContext(c)
	if err != nil {
		return c.Redirect(http.StatusNotFound, "/hubs")
	}

	pageSlug := c.Param("pageSlug")

	page, err := h.Ent.Page.Query().
		Where(entPage.HasHubWith(hub.ID(contextHub.ID))).
		Where(page.Slug(pageSlug)).
		Only(ctx)

	if err != nil {
		return c.Redirect(http.StatusNotFound, "/hubs")
	}

	payload := new(PagePayload)

	if err := extensions.BindAndValidate(c, payload); err != nil {
		feedback := partial.FormFeedbackFromErr("Edit Page", err)
		return extensions.Render(c, http.StatusOK, editForm(contextHub.Slug, *payload, feedback))
	}

	metaPayload := new(partial.MetaDescriptionPayload)
	if err := extensions.BindAndValidate(c, metaPayload); err != nil {
		feedback := partial.FormFeedbackFromErr("Edit Page", err)
		return extensions.Render(c, http.StatusOK, editForm(contextHub.Slug, *payload, feedback))
	}

	payload.MetaData = metaPayload.AsModel()

	query := h.Ent.Page.UpdateOne(page).
		SetName(payload.Name).
		SetPageDescription(payload.Description).
		SetPageContentJSON(payload.ContentJSON).
		SetMetaDescription(payload.MetaData).
		SetSlug(payload.PageSlug)

	if payload.ContentHTML != "" {
		query.SetPageContentHTML(payload.ContentHTML)
	}

	updatedPage, err := query.Save(ctx)

	if err != nil {
		feedback := partial.FormFeedbackFromErr("Edit Page", err)
		return extensions.Render(c, http.StatusOK, editForm(contextHub.Slug, *payload, feedback))
	}

	extensions.HTMXAppendSuccessToast(c, "Page updated successfully")
	extensions.HTMXCloseModal(c)
	extensions.HTMXAppendPrelineRefresh(c)
	extensions.HTMXAppendEventsAfterSwap(c, map[string]interface{}{
		"grapejs-reinit": map[string]interface{}{
			"id": updatedPage.Slug,
		},
	})

	payload.UpdateFromEntity(updatedPage)

	return extensions.Render(c, http.StatusOK, editForm(contextHub.Slug, *payload, nil))
}

func (h Handlers) PageStatusGet(c echo.Context) error {
	ctx := c.Request().Context()

	_, err := context.GetUserFromContext(c)
	if err != nil {
		return c.NoContent(http.StatusForbidden)
	}

	contextHub, _ := context.GetHubFromContext(c)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	pageSlug := c.Param("pageSlug")

	page, err := h.Ent.Page.Query().
		Where(entPage.HasHubWith(hub.ID(contextHub.ID))).
		Where(page.Slug(pageSlug)).
		Only(ctx)

	if err != nil {
		return c.Redirect(http.StatusNotFound, "/hubs")
	}

	payload := EditPageStatusPayload{
		Status:   page.Status,
		PageSlug: page.Slug,
		HubSlug:  contextHub.Slug,
	}

	return extensions.Render(c, http.StatusOK, PageStatusForm(payload))
}

func (h Handlers) PageStatusPost(c echo.Context) error {
	ctx := c.Request().Context()

	_, err := context.GetUserFromContext(c)
	if err != nil {
		return c.NoContent(http.StatusForbidden)
	}

	contextHub, _ := context.GetHubFromContext(c)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	pageSlug := c.Param("pageSlug")

	page, err := h.Ent.Page.Query().
		Where(entPage.HasHubWith(hub.ID(contextHub.ID))).
		Where(page.Slug(pageSlug)).
		Only(ctx)

	if err != nil {
		return c.Redirect(http.StatusNotFound, "/hubs")
	}

	payload := new(EditPageStatusPayload)

	if err := extensions.BindAndValidate(c, payload); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	hub, _ := context.GetHubFromContext(c)

	payload.HubSlug = hub.Slug
	payload.PageSlug = page.Slug

	page, err = h.Ent.Page.UpdateOne(page).
		SetStatus(payload.Status).
		Save(c.Request().Context())

	if err != nil {
		h.Logger.Error("Error updating page", "error", err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	extensions.HTMXAppendSuccessToast(c, "Link Status Updated Successfully")
	extensions.HTMXAppendPrelineRefresh(c)

	return extensions.Render(c, http.StatusOK, PageStatusForm(*payload))
}
