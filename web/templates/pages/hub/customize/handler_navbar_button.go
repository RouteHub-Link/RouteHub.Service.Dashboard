package customize

import (
	"net/http"
	"strings"

	"RouteHub.Service.Dashboard/ent/schema/types"
	"RouteHub.Service.Dashboard/web/context"
	"RouteHub.Service.Dashboard/web/extensions"
	"RouteHub.Service.Dashboard/web/templates/pages/partial"
	"github.com/labstack/echo/v4"
)

func (h Handlers) NavbarButtonEditFormGetHandler(c echo.Context) error {
	hub, _ := context.GetHubFromContext(c)
	title := "Navbar Button Form"

	payload := NavbarItemButtonFormPayload{
		HubSlug: hub.Slug,
	}

	itemID := c.Param("itemID")
	if itemID != "" {
		item, err := GetButtonByIndex(&hub.HubDetails.NavbarDescription, itemID)
		if err != nil {
			msg := "Error finding item"
			feedback := partial.FormFeedback("error", &title, &msg)
			return extensions.Render(c, http.StatusOK, NavbarItemButtonForm(payload, feedback))
		}

		payload.ID = itemID
		payload.Name = item.Text
		payload.URL = item.URL
		payload.Target = item.Target
		payload.ColorClass = item.ColorClass

	}

	return extensions.Render(c, http.StatusOK, NavbarItemButtonForm(payload, nil))
}

func (h Handlers) NavbarButtonDeleteFormGetHandler(c echo.Context) error {
	hub, _ := context.GetHubFromContext(c)
	title := "Navbar Button Form"

	payload := NavbarItemButtonFormPayload{
		HubSlug: hub.Slug,
	}

	itemID := c.Param("itemID")
	if itemID == "" {
		msg := "Error finding item"
		feedback := partial.FormFeedback("error", &title, &msg)
		return extensions.Render(c, http.StatusOK, NavbarItemButtonDeleteForm(payload, feedback))
	}

	item, err := GetButtonByIndex(&hub.HubDetails.NavbarDescription, itemID)
	if err != nil {
		msg := "Error finding item"
		feedback := partial.FormFeedback("error", &title, &msg)
		return extensions.Render(c, http.StatusOK, NavbarItemButtonDeleteForm(payload, feedback))
	}

	payload.ID = itemID
	payload.Name = item.Text
	payload.URL = item.URL
	payload.Target = item.Target
	payload.ColorClass = item.ColorClass

	return extensions.Render(c, http.StatusOK, NavbarItemButtonDeleteForm(payload, nil))
}

func (h Handlers) NavbarButtonDeleteFormPostHandler(c echo.Context) error {
	hub, _ := context.GetHubFromContext(c)
	title := "Navbar Button Form"

	payload := NavbarItemButtonFormPayload{
		HubSlug: hub.Slug,
	}

	itemID := c.Param("itemID")
	if itemID == "" {
		msg := "Error finding item"
		feedback := partial.FormFeedback("error", &title, &msg)
		return extensions.Render(c, http.StatusOK, NavbarItemButtonForm(payload, feedback))
	}

	_, err := GetButtonByIndex(&hub.HubDetails.NavbarDescription, itemID)
	if err != nil {
		msg := "Error finding item"
		feedback := partial.FormFeedback("error", &title, &msg)
		return extensions.Render(c, http.StatusOK, NavbarItemButtonForm(payload, feedback))
	}

	err = DeleteButtonByIndex(&hub.HubDetails.NavbarDescription, itemID)

	if err != nil {
		msg := "Error deleting item"
		feedback := partial.FormFeedback("error", &title, &msg)
		return extensions.Render(c, http.StatusOK, NavbarItemButtonForm(payload, feedback))
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

func (h Handlers) NavbarButtonEditFormPostHandler(c echo.Context) error {
	hub, _ := context.GetHubFromContext(c)
	title := "Navbar Button Form"

	payload := new(NavbarItemButtonFormPayload)
	if err := extensions.BindAndValidate(c, payload); err != nil {
		msg := strings.Join([]string{"Error Validating Data", err.Error()}, " ")
		feedback := partial.FormFeedback("error", &title, &msg)
		return extensions.Render(c, http.StatusOK, NavbarItemButtonForm(*payload, feedback))
	}

	itemID := c.Param("itemID")
	item := new(types.NavbarButton)
	isNew := true

	if itemID != "" && itemID != "nav-root" {
		gettedItem, err := GetButtonByIndex(&hub.HubDetails.NavbarDescription, itemID)

		if err != nil {
			msg := "Error finding item"
			feedback := partial.FormFeedback("error", &title, &msg)
			return extensions.Render(c, http.StatusOK, NavbarItemButtonForm(*payload, feedback))
		}

		item = gettedItem
		isNew = false
	}

	item.Text = payload.Name
	item.URL = payload.URL
	item.Target = payload.Target
	item.Icon = payload.Icon
	item.ColorClass = payload.ColorClass

	switch isNew {
	case true:
		AppendButton(&hub.HubDetails.NavbarDescription, *item)
	case false:
		err := ReplaceButtonByIndex(&hub.HubDetails.NavbarDescription, itemID, *item)

		if err != nil {
			msg := "Error updating item"
			feedback := partial.FormFeedback("error", &title, &msg)
			return extensions.Render(c, http.StatusOK, NavbarItemButtonForm(*payload, feedback))
		}
	}

	hub, err := h.Ent.Hub.UpdateOne(hub).SetHubDetails(hub.HubDetails).Save(c.Request().Context())
	if err != nil {
		msg := "Error updating item"
		feedback := partial.FormFeedback("error", &title, &msg)
		return extensions.Render(c, http.StatusOK, NavbarItemButtonForm(*payload, feedback))
	}

	extensions.HTMXAppendSuccessToast(c, "Item Updated Successfully")
	extensions.HTMXAppendPrelineRefresh(c)
	extensions.HTMXAppendEventsAfterSwap(c, map[string]interface{}{
		"navbarItemUpdated": "",
	})
	extensions.HTMXCloseModal(c)

	return nil
}
