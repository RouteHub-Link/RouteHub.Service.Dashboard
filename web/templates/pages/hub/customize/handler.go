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

func (h Handlers) CustomizeGetHandler(c echo.Context) error {

	userInfo, err := context.GetUserFromContext(c)
	if err != nil {
		return err
	}
	hub, err := context.GetHubFromContext(c)

	return extensions.Render(c, http.StatusOK, index(userInfo, hub))
}

func (h Handlers) MetaPageGet(c echo.Context) error {
	hub, _ := context.GetHubFromContext(c)
	userInfo, _ := context.GetUserFromContext(c)

	hubMeta := hub.HubDetails.MetaDescription

	MetaPayload := MetaPayload{
		Title:         hubMeta.Title,
		Description:   hubMeta.Description,
		OGDescription: hubMeta.OGDescription,
		OGTitle:       hubMeta.OGTitle,
		Locale:        hubMeta.Locale,
		TracingScript: hubMeta.TracingScript,
		FavIcon:       utils.LinkToS3Path(hubMeta.FavIcon),
	}

	extensions.HTMXAppendPrelineRefresh(c)
	return extensions.Render(c, http.StatusOK, MetaPage(userInfo, hub, MetaPayload, hub.Slug, nil))
}

func (h Handlers) MetaPartialGet(c echo.Context) error {
	hub, _ := context.GetHubFromContext(c)

	hubMeta := hub.HubDetails.MetaDescription

	MetaPayload := MetaPayload{
		Title:         hubMeta.Title,
		Description:   hubMeta.Description,
		OGDescription: hubMeta.OGDescription,
		OGTitle:       hubMeta.OGTitle,
		Locale:        hubMeta.Locale,
		TracingScript: hubMeta.TracingScript,
		FavIcon:       utils.LinkToS3Path(hubMeta.FavIcon),
	}

	extensions.HTMXAppendPrelineRefresh(c)
	return extensions.Render(c, http.StatusOK, metaForm(MetaPayload, hub.Slug, nil))
}

func (h Handlers) MetaPartialPost(c echo.Context) error {
	hub, _ := context.GetHubFromContext(c)
	title := "Meta Data Form"

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
	hubDetails.MetaDescription.TracingScript = payload.TracingScript

	bucketPath := strings.Join([]string{"hubs", hub.Slug}, "/")

	if err := extensions.ProcessFileFromEchoContext(c, &payload.FavIcon, "meta_description_favicon", bucketPath, "favicon"); err != nil {
		feedback := partial.FormFeedbackFromErr(title, fmt.Errorf("Error Processing File: %w", err))
		return extensions.Render(c, http.StatusOK, metaForm(*payload, hub.Slug, feedback))
	}

	hubDetails.MetaDescription.FavIcon = payload.FavIcon

	if _, err := h.Ent.Hub.
		UpdateOne(hub).
		SetHubDetails(hubDetails).
		Save(c.Request().Context()); err != nil {

		feedback := partial.FormFeedbackFromErr(title, err)
		return extensions.Render(c, http.StatusOK, metaForm(*payload, hub.Slug, feedback))
	}

	extensions.HTMXAppendSuccessToast(c, "Meta Updated Successfully")
	extensions.HTMXAppendPrelineRefresh(c)

	payload.FavIcon = utils.LinkToS3Path(payload.FavIcon)

	return extensions.Render(c, http.StatusOK, metaForm(*payload, hub.Slug, nil))
}

type NavbarCustomizePayload struct {
	BrandName string `json:"brand_name" form:"brand_name" validate:"required"`
	BrandURL  string `json:"brand_url" form:"brand_url" validate:"required,url"`
}

func (h Handlers) NavbarPageGet(c echo.Context) error {
	hub, _ := context.GetHubFromContext(c)
	userInfo, _ := context.GetUserFromContext(c)

	return extensions.Render(c, http.StatusOK, NavbarPage(userInfo, hub))
}

func (h Handlers) NavbarPartialGet(c echo.Context) error {
	hub, _ := context.GetHubFromContext(c)

	payload := NavbarCustomizePayload{
		BrandName: hub.HubDetails.NavbarDescription.BrandName,
		BrandURL:  hub.HubDetails.NavbarDescription.BrandURL,
	}

	return extensions.Render(c, http.StatusOK, navbarForm(payload, hub.HubDetails.NavbarDescription, hub.Slug, nil))
}

func (h Handlers) NavbarPartialPost(c echo.Context) error {
	hub, _ := context.GetHubFromContext(c)
	title := "Navbar Form"

	payload := new(NavbarCustomizePayload)
	if err := extensions.BindAndValidate(c, payload); err != nil {
		feedback := partial.FormFeedback("error", title, err.Error())
		return extensions.Render(c, http.StatusOK, navbarForm(*payload, hub.HubDetails.NavbarDescription, hub.Slug, feedback))
	}

	hubDetails := hub.HubDetails
	hubDetails.NavbarDescription.BrandName = payload.BrandName
	hubDetails.NavbarDescription.BrandURL = payload.BrandURL

	if _, err := h.Ent.Hub.
		UpdateOne(hub).
		SetHubDetails(hubDetails).
		Save(c.Request().Context()); err != nil {
		feedback := partial.FormFeedbackFromErr(title, err)
		return extensions.Render(c, http.StatusOK, navbarForm(*payload, hub.HubDetails.NavbarDescription, hub.Slug, feedback))
	}

	extensions.HTMXAppendSuccessToast(c, "Navbar Updated Successfully")
	extensions.HTMXAppendPrelineRefresh(c)
	extensions.HTMXAppendEventsAfterSwap(c, map[string]interface{}{
		"navbarUpdated": "",
	})

	return nil
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
		feedback := partial.FormFeedback("error", title, msg)

		return extensions.Render(c, http.StatusOK, NavbarItemForm(*payload, feedback))
	}

	// itemId
	itemID := c.Param("itemID")
	if itemID == "" {
		feedback := partial.FormFeedback("error", title, "Error finding item")
		return extensions.Render(c, http.StatusOK, NavbarItemForm(*payload, feedback))
	}

	item, err := GetItemByIndex(&hub.HubDetails.NavbarDescription, itemID)
	if err != nil {
		feedback := partial.FormFeedback("error", title, "Error finding item")
		return extensions.Render(c, http.StatusOK, NavbarItemForm(*payload, feedback))
	}

	item.Text = payload.Name
	item.URL = payload.URL
	item.Target = payload.Target
	item.Icon = payload.Icon

	err = ReplaceItemByIndex(&hub.HubDetails.NavbarDescription, itemID, *item)

	if err != nil {
		feedback := partial.FormFeedback("error", title, "Error updating item")
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
		feedback := partial.FormFeedback("error", "", message)
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
		feedback := partial.FormFeedback("error", title, msg)
		return extensions.Render(c, http.StatusOK, NavbarItemForm(*payload, feedback))
	}

	itemID := c.Param("itemID")
	if itemID == "" {
		feedback := partial.FormFeedback("error", title, "Error finding item")
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
			feedback := partial.FormFeedback("error", title, "Error finding item")
			return extensions.Render(c, http.StatusOK, NavbarItemForm(*payload, feedback))
		}
	}

	hub, err := h.Ent.Hub.UpdateOne(hub).SetHubDetails(hub.HubDetails).Save(c.Request().Context())
	if err != nil {
		feedback := partial.FormFeedback("error", title, "Error updating item")
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
	title := "Navbar Item Form"

	payload := NavbarItemFormPayload{
		HubSlug: hub.Slug,
	}

	itemID := c.Param("itemID")
	if itemID == "" {
		feedback := partial.FormFeedback("error", title, "Error finding item")
		return extensions.Render(c, http.StatusOK, NavbarItemDeleteForm(payload, feedback))
	}

	_, err := GetItemByIndex(&hub.HubDetails.NavbarDescription, itemID)
	if err != nil {
		feedback := partial.FormFeedback("error", title, "Error finding item")
		return extensions.Render(c, http.StatusOK, NavbarItemDeleteForm(payload, feedback))
	}

	err = DeleteItemByIndex(&hub.HubDetails.NavbarDescription, itemID)

	if err != nil {
		feedback := partial.FormFeedback("error", title, "Error deleting item")
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
