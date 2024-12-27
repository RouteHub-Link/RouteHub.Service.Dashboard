package customize

import (
	"fmt"
	"net/http"

	"RouteHub.Service.Dashboard/ent/schema/types"
	"RouteHub.Service.Dashboard/web/context"
	"RouteHub.Service.Dashboard/web/extensions"
	"RouteHub.Service.Dashboard/web/templates/pages/partial"
	"github.com/labstack/echo/v4"
)

type FooterPayload struct {
	ShowRouteHubBranding bool `json:"show_routehub_branding" form:"show_routehub_branding"`
}

func (h Handlers) FooterPageGet(c echo.Context) error {
	hub, _ := context.GetHubFromContext(c)
	userInfo, _ := context.GetUserFromContext(c)

	return extensions.Render(c, http.StatusOK, FooterPage(userInfo, hub))
}

func (h Handlers) FooterPartialGet(c echo.Context) error {
	hub, _ := context.GetHubFromContext(c)

	payload := FooterPayload{
		ShowRouteHubBranding: hub.HubDetails.FooterDescription.ShowRouteHubBranding,
	}

	extensions.HTMXAppendPrelineRefresh(c)

	return extensions.Render(c, http.StatusOK, footerForm(payload, hub.HubDetails.FooterDescription, hub.Slug, nil))
}

func (h Handlers) FooterPartialPost(c echo.Context) error {
	hub, _ := context.GetHubFromContext(c)
	title := "Footer Customization"

	payload := new(FooterPayload)
	if err := extensions.BindAndValidate(c, payload); err != nil {
		payload = &FooterPayload{
			ShowRouteHubBranding: hub.HubDetails.FooterDescription.ShowRouteHubBranding,
		}
		feedback := partial.FormFeedback("error", title, err.Error())
		return extensions.Render(c, http.StatusBadRequest, footerForm(*payload, hub.HubDetails.FooterDescription, hub.Slug, feedback))
	}

	hubDetails := hub.HubDetails
	hubDetails.FooterDescription.ShowRouteHubBranding = payload.ShowRouteHubBranding

	if _, err := h.Ent.Hub.UpdateOne(hub).SetHubDetails(hubDetails).Save(c.Request().Context()); err != nil {
		payload = &FooterPayload{
			ShowRouteHubBranding: hub.HubDetails.FooterDescription.ShowRouteHubBranding,
		}
		feedback := partial.FormFeedback("error", title, err.Error())
		return extensions.Render(c, http.StatusBadRequest, footerForm(*payload, hub.HubDetails.FooterDescription, hub.Slug, feedback))
	}

	extensions.HTMXAppendSuccessToast(c, "Footer updated successfully")
	extensions.HTMXAppendEventsAfterSwap(c, map[string]interface{}{
		"footerUpdated": "",
	})
	extensions.HTMXAppendPrelineRefresh(c)
	extensions.HTMXCloseModal(c)

	feedback := partial.FormFeedback("success", title, "Footer updated successfully")

	return extensions.Render(c, http.StatusOK, footerForm(*payload, hub.HubDetails.FooterDescription, hub.Slug, feedback))
}

func (h Handlers) FooterSocialMediaContainerGet(c echo.Context) error {
	hub, _ := context.GetHubFromContext(c)
	socialMediaContainer := types.SocialMediaContainer{}
	if hub.HubDetails.FooterDescription.SocialMediaContainer != nil {
		socialMediaContainer = *hub.HubDetails.FooterDescription.SocialMediaContainer
	}

	return extensions.Render(c, http.StatusOK, footerSocialMediaContainer(&socialMediaContainer, hub.Slug))
}

func (h Handlers) FooterSocialMediaLinksGet(c echo.Context) error {
	hub, _ := context.GetHubFromContext(c)
	socialMediaLinks := []types.ASocialMedia{}
	if hub.HubDetails.FooterDescription.SocialMediaContainer == nil {
		return c.JSON(http.StatusOK, socialMediaLinks)
	}

	_socialMediaLinks := hub.HubDetails.FooterDescription.SocialMediaContainer.SocialMediaLinks
	if _socialMediaLinks != nil {
		socialMediaLinks = *_socialMediaLinks
	}
	return c.JSON(http.StatusOK, socialMediaLinks)
}

type SocialMediaPayload struct {
	SocialMediaLinks []types.ASocialMedia `json:"socialMediaLinks"`
}

func (h Handlers) FooterSocialMediaLinksPost(c echo.Context) error {
	hub, _ := context.GetHubFromContext(c)

	payload := new(SocialMediaPayload)
	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if hub.HubDetails.FooterDescription.SocialMediaContainer == nil {
		hub.HubDetails.FooterDescription.SocialMediaContainer = &types.SocialMediaContainer{}
	}

	hub.HubDetails.FooterDescription.SocialMediaContainer.SocialMediaLinks = &payload.SocialMediaLinks
	if _, err := h.Ent.Hub.UpdateOne(hub).SetHubDetails(hub.HubDetails).Save(c.Request().Context()); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	extensions.HTMXAppendSuccessToast(c, "Branding HTML updated successfully")
	extensions.HTMXAppendEventsAfterSwap(c, map[string]interface{}{
		"footerItemUpdated": "",
	})
	extensions.HTMXCloseModal(c)

	return c.NoContent(http.StatusOK)
}

func (h Handlers) FooterShadow(c echo.Context) error {
	hub, _ := context.GetHubFromContext(c)
	shadowID := "footer"
	resultHTML := ""

	if hub.HubDetails.FooterDescription.CompanyBrandingHTML != "" {
		resultHTML = hub.HubDetails.FooterDescription.CompanyBrandingHTML
	}

	return extensions.Render(c, http.StatusOK, partial.ShadowHub(FooterShadow(&hub.HubDetails.FooterDescription, resultHTML), shadowID))
}

func (h Handlers) FooterBrandingHTMLGet(c echo.Context) error {
	hub, _ := context.GetHubFromContext(c)
	return extensions.Render(c, http.StatusOK, BrandingHTML(hub.HubDetails.FooterDescription, hub.Slug))
}

type BrandingHTMLPayload struct {
	CompanyBrandingHTML string `json:"companyBrandingHTML" form:"companyBrandingHTML" validate:"required"`
}

func (h Handlers) FooterBrandingHTMLPost(c echo.Context) error {
	hub, _ := context.GetHubFromContext(c)

	payload := new(BrandingHTMLPayload)
	if err := extensions.BindAndValidate(c, payload); err != nil {
		extensions.HTMXAppendErrorToast(c, fmt.Sprintf("Error: %v", err))
		return err
	}

	hub.HubDetails.FooterDescription.CompanyBrandingHTML = payload.CompanyBrandingHTML
	h.Ent.Hub.UpdateOne(hub).SetHubDetails(hub.HubDetails).Save(c.Request().Context())

	extensions.HTMXAppendSuccessToast(c, "Branding HTML updated successfully")
	extensions.HTMXCloseModal(c)
	extensions.HTMXAppendEventsAfterSwap(c, map[string]interface{}{
		"footerItemUpdated": "",
	})

	return c.NoContent(http.StatusOK)
}
