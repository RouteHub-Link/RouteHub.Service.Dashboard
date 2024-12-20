package customize

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"RouteHub.Service.Dashboard/ent"
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

func (h Handlers) CustomizeGetHandler(c echo.Context) error {

	userInfo, err := context.GetUserFromContext(c)
	if err != nil {
		return err
	}
	hub, err := context.GetHubFromContext(c)

	return extensions.Render(c, http.StatusOK, index(userInfo, hub))
}

func (h Handlers) MetaGet(c echo.Context) error {
	hub, _ := context.GetHubFromContext(c)

	hubMeta := hub.HubDetails.MetaDescription

	MetaPayload := MetaPayload{
		Title:         hubMeta.Title,
		Description:   hubMeta.Description,
		OGDescription: hubMeta.OGDescription,
		OGTitle:       hubMeta.OGTitle,
		Locale:        hubMeta.Locale,
	}

	extensions.HTMXAppendPrelineRefresh(c)
	return extensions.Render(c, http.StatusOK, meta(MetaPayload))
}

func (h Handlers) MetaPost(c echo.Context) error {
	hub, _ := context.GetHubFromContext(c)

	payload := new(MetaPayload)
	if err := extensions.BindAndValidate(c, payload); err != nil {
		return err
	}

	hubDetails := hub.HubDetails
	hubDetails.MetaDescription.Title = payload.Title
	hubDetails.MetaDescription.Description = payload.Description
	hubDetails.MetaDescription.OGDescription = payload.OGDescription
	hubDetails.MetaDescription.OGTitle = payload.OGTitle
	hubDetails.MetaDescription.Locale = payload.Locale

	if _, err := h.Ent.Hub.
		UpdateOne(hub).
		SetHubDetails(hubDetails).
		Save(c.Request().Context()); err != nil {

		h.Logger.Error("Error updating hub", "error", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	extensions.HTMXAppendSuccessToast(c, "Meta Updated Successfully")
	extensions.HTMXAppendPrelineRefresh(c)

	return c.NoContent(http.StatusOK)
}

func (h Handlers) NavbarGet(c echo.Context) error {
	hub, _ := context.GetHubFromContext(c)
	//mockMetaDescription := types.MetaDescription{Title: "RouteHub", Description: "RouteHub is a platform that allows you to create, share, and discover routes for your favorite activities."}

	//brandImg := types.ImageDescription{SRC: "https://avatars.githubusercontent.com/u/153122518?s=250", Alt: "RouteHub", Width: "30", Height: "30"}
	//navbarItems := []types.NavbarItem{{Text: "Home", URL: "/", Target: "_self", Icon: "home"}, {Text: "About", URL: "/about", Target: "_self", Icon: "info", Dropdown: &[]types.NavbarItem{{Text: "Contact", URL: "/contact", Target: "_self", Icon: "contact_mail"}}}}
	//navbarEndButtons := []types.NavbarButton{{Text: "Login", URL: "/login", Target: "_self", ColorClass: "is-secondary"}, {Text: "Sign Up", URL: "/signup", Target: "_self", ColorClass: "is-primary"}}

	//navbarDescription := types.NavbarDescription{BrandName: "RouteHub", BrandURL: "https://routehub.link", BrandImg: &brandImg, StartItems: &navbarItems, EndButtons: &navbarEndButtons}

	//socialMediaList := []types.ASocialMedia{{Icon: "facebook", Link: "https://www.facebook.com", Target: "_blank"}, {Icon: "twitter", Link: "https://www.twitter.com", Target: "_blank"}, {Icon: "instagram", Link: "https://www.instagram.com", Target: "_blank"}, {Icon: "linkedin", Link: "https://www.linkedin.com", Target: "_blank"}}
	//socialMediaContainer := types.SocialMediaContainer{SocialMediaLinks: &socialMediaList, SocialMediaPeddingClass: "pt-5", SocialMediaSizeClass: "is-medium", SocialMediaColorClass: "has-text-white"}

	//footerDescription := types.FooterDescription{ShowRouteHubBranding: true, CompanyBrandingHTML: "<strong>X Company</strong> <a href=''> X Company</a> Has Rights of this site since 1111</strong>", SocialMediaContainer: &socialMediaContainer}

	//hub.HubDetails.FooterDescription = footerDescription
	//hub.HubDetails.NavbarDescription = navbarDescription
	//hub.HubDetails.MetaDescription = mockMetaDescription

	//hub, _ = h.Ent.Hub.UpdateOne(hub).SetHubDetails(hub.HubDetails).Save(c.Request().Context())

	return extensions.Render(c, http.StatusOK, navbar(hub.HubDetails.NavbarDescription, hub.Slug))
}

func (h Handlers) NavbarItemEditFormGet(c echo.Context) error {
	fmt.Println("NavbarEditItemFormGetHandler Called")
	hub, _ := context.GetHubFromContext(c)

	payload := new(NavbarItemFormPayload)

	// itemId
	itemID := c.Param("itemID")
	fmt.Printf("ItemID : %s", itemID)
	if itemID != "" {
		item, err := GetItemByIndex(&hub.HubDetails.NavbarDescription, itemID)
		fmt.Printf("item : %+v", item)
		fmt.Printf("err : %v", err)

		if err != nil {
			return c.NoContent(http.StatusNotFound)
		}
		payload = &NavbarItemFormPayload{
			HubSlug: hub.Slug,
			Name:    item.Text,
			URL:     item.URL,
			Target:  item.Target,
			Icon:    item.Icon,
		}
	}

	if payload == nil {
		return c.NoContent(http.StatusNotFound)
	}

	payload.ID = itemID

	return extensions.Render(c, http.StatusOK, NavbarItemForm(*payload, nil))
}

func (h Handlers) NavbarItemEditFormPost(c echo.Context) error {
	hub, _ := context.GetHubFromContext(c)
	title := "Navbar Item Form"

	payload := new(NavbarItemFormPayload)
	if err := extensions.BindAndValidate(c, payload); err != nil {
		msg := strings.Join([]string{"Error Validating Data", err.Error()}, " ")
		feedback := partial.FormFeedback("error", &title, &msg)
		return extensions.Render(c, http.StatusOK, NavbarItemForm(*payload, feedback))
	}

	// itemId
	itemID := c.Param("itemID")
	if itemID == "" {
		msg := "Error finding item"
		feedback := partial.FormFeedback("error", &title, &msg)
		return extensions.Render(c, http.StatusOK, NavbarItemForm(*payload, feedback))
	}

	item, err := GetItemByIndex(&hub.HubDetails.NavbarDescription, itemID)
	if err != nil {
		msg := "Error finding item"
		feedback := partial.FormFeedback("error", &title, &msg)
		return extensions.Render(c, http.StatusOK, NavbarItemForm(*payload, feedback))
	}

	item.Text = payload.Name
	item.URL = payload.URL
	item.Target = payload.Target
	item.Icon = payload.Icon

	err = ReplaceItemByIndex(&hub.HubDetails.NavbarDescription, itemID, *item)

	if err != nil {
		msg := "Error updating item"
		feedback := partial.FormFeedback("error", &title, &msg)
		return extensions.Render(c, http.StatusOK, NavbarItemForm(*payload, feedback))
	}

	hub, err = h.Ent.Hub.UpdateOne(hub).SetHubDetails(hub.HubDetails).Save(c.Request().Context())

	extensions.HTMXAppendSuccessToast(c, "Item Updated Successfully")
	extensions.HTMXAppendPrelineRefresh(c)
	extensions.HTMXAppendEventsAfterSwap(c, map[string]interface{}{
		"navbarItemUpdated": "",
	})
	extensions.HTMXCloseModal(c)

	return nil
}

func (h Handlers) NavbarItemAddFormGet(c echo.Context) error {
	hub, _ := context.GetHubFromContext(c)

	itemID := c.Param("itemID")
	payload := NavbarItemFormPayload{
		HubSlug: hub.Slug,
		ID:      itemID,
	}

	if itemID == "nav-root" {
		return extensions.Render(c, http.StatusOK, NavbarItemAddForm(payload, nil))
	}

	_, err := GetItemByIndex(&hub.HubDetails.NavbarDescription, itemID)
	if err != nil {
		message := fmt.Sprintf("Error finding item with id %s", itemID)
		feedback := partial.FormFeedback("error", nil, &message)
		return extensions.Render(c, http.StatusOK, NavbarItemAddForm(NavbarItemFormPayload{}, feedback))
	}

	return extensions.Render(c, http.StatusOK, NavbarItemAddForm(payload, nil))
}

func (h Handlers) NavbarItemAddFormPost(c echo.Context) error {
	hub, _ := context.GetHubFromContext(c)
	title := "Navbar Item Form"

	payload := new(NavbarItemFormPayload)
	if err := extensions.BindAndValidate(c, payload); err != nil {
		msg := strings.Join([]string{"Error Validating Data", err.Error()}, " ")
		feedback := partial.FormFeedback("error", &title, &msg)
		return extensions.Render(c, http.StatusOK, NavbarItemForm(*payload, feedback))
	}

	itemID := c.Param("itemID")
	if itemID == "" {
		msg := "Error finding item"
		feedback := partial.FormFeedback("error", &title, &msg)
		return extensions.Render(c, http.StatusOK, NavbarItemForm(*payload, feedback))
	}

	newItem := types.NavbarItem{
		Text:   payload.Name,
		URL:    payload.URL,
		Target: payload.Target,
		Icon:   payload.Icon,
	}

	if itemID == "nav-root" {
		startItems := hub.HubDetails.NavbarDescription.StartItems
		if startItems == nil {
			startItems = &[]types.NavbarItem{}
		}
		*startItems = append(*startItems, newItem)
		hub.HubDetails.NavbarDescription.StartItems = startItems

	} else {
		_, err := navbarItemAppendToItem(&hub.HubDetails.NavbarDescription, itemID, newItem)

		if err != nil {
			msg := "Error finding item"
			feedback := partial.FormFeedback("error", &title, &msg)
			return extensions.Render(c, http.StatusOK, NavbarItemForm(*payload, feedback))
		}
	}

	hub, err := h.Ent.Hub.UpdateOne(hub).SetHubDetails(hub.HubDetails).Save(c.Request().Context())
	if err != nil {
		msg := "Error updating item"
		feedback := partial.FormFeedback("error", &title, &msg)
		return extensions.Render(c, http.StatusOK, NavbarItemForm(*payload, feedback))
	}

	extensions.HTMXAppendSuccessToast(c, "Item Updated Successfully")
	extensions.HTMXAppendPrelineRefresh(c)
	extensions.HTMXAppendEventsAfterSwap(c, map[string]interface{}{
		"navbarItemUpdated": "",
	})
	extensions.HTMXCloseModal(c)

	return nil
}

func navbarItemAppendToItem(nav *types.NavbarDescription, indexPath string, newItem types.NavbarItem) (*types.NavbarItem, error) {
	item, err := GetItemByIndex(nav, indexPath)

	if err != nil {
		msg := "Error finding item"
		return nil, errors.New(msg)
	}

	if item.Dropdown == nil {
		item.Dropdown = &[]types.NavbarItem{}
	}

	*item.Dropdown = append(*item.Dropdown, newItem)

	err = ReplaceItemByIndex(nav, indexPath, *item)

	if err != nil {
		msg := "Error updating item"
		return nil, errors.New(msg)
	}

	return item, nil
}

func (h Handlers) NavbarItemDeleteFormGet(c echo.Context) error {
	fmt.Println("NavbarEditItemFormGetHandler Called")
	hub, _ := context.GetHubFromContext(c)

	payload := new(NavbarItemFormPayload)

	// itemId
	itemID := c.Param("itemID")
	if itemID != "" {
		item, err := GetItemByIndex(&hub.HubDetails.NavbarDescription, itemID)

		if err != nil {
			return err
		}

		payload = &NavbarItemFormPayload{
			HubSlug: hub.Slug,
			Name:    item.Text,
			URL:     item.URL,
			Target:  item.Target,
			Icon:    item.Icon,
		}
	}

	if payload == nil {
		return c.NoContent(http.StatusNotFound)
	}

	payload.ID = itemID

	return extensions.Render(c, http.StatusOK, NavbarItemDeleteForm(*payload, nil))
}

func (h Handlers) NavbarItemDeletePost(c echo.Context) error {
	hub, _ := context.GetHubFromContext(c)

	payload := NavbarItemFormPayload{
		HubSlug: hub.Slug,
	}

	itemID := c.Param("itemID")
	if itemID == "" {
		msg := "Error finding item"
		feedback := partial.FormFeedback("error", nil, &msg)
		return extensions.Render(c, http.StatusOK, NavbarItemDeleteForm(payload, feedback))
	}

	_, err := GetItemByIndex(&hub.HubDetails.NavbarDescription, itemID)
	if err != nil {
		msg := "Error finding item"
		feedback := partial.FormFeedback("error", nil, &msg)
		return extensions.Render(c, http.StatusOK, NavbarItemDeleteForm(payload, feedback))
	}

	err = DeleteItemByIndex(&hub.HubDetails.NavbarDescription, itemID)

	if err != nil {
		msg := "Error deleting item"
		feedback := partial.FormFeedback("error", nil, &msg)
		return extensions.Render(c, http.StatusOK, NavbarItemDeleteForm(payload, feedback))
	}

	hub, err = h.Ent.Hub.UpdateOne(hub).SetHubDetails(hub.HubDetails).Save(c.Request().Context())

	extensions.HTMXAppendSuccessToast(c, "Item Deleted Successfully")
	extensions.HTMXAppendPrelineRefresh(c)
	extensions.HTMXAppendEventsAfterSwap(c, map[string]interface{}{
		"navbarItemUpdated": "",
	})
	extensions.HTMXCloseModal(c)

	return nil
}
