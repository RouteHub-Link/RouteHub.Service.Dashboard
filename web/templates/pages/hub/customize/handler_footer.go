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

func (h Handlers) FooterGet(c echo.Context) error {
	hub, _ := context.GetHubFromContext(c)

	return extensions.Render(c, http.StatusOK, footer(hub.HubDetails.FooterDescription, hub.Slug))
}

func (h Handlers) FooterSocialMediaContainerGet(c echo.Context) error {
	hub, _ := context.GetHubFromContext(c)
	return extensions.Render(c, http.StatusOK, footerSocialMediaContainer(hub.HubDetails.FooterDescription.SocialMediaContainer, hub.Slug))
}

func (h Handlers) FooterSocialMediaLinksGet(c echo.Context) error {
	hub, _ := context.GetHubFromContext(c)
	return c.JSON(http.StatusOK, hub.HubDetails.FooterDescription.SocialMediaContainer.SocialMediaLinks)
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

	hub.HubDetails.FooterDescription.SocialMediaContainer.SocialMediaLinks = &payload.SocialMediaLinks
	h.Ent.Hub.UpdateOne(hub).SetHubDetails(hub.HubDetails).Save(c.Request().Context())

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
