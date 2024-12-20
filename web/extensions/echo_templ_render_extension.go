package extensions

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

/*
Response Headers
htmx supports some htmx-specific response headers:

HX-Location - allows you to do a client-side redirect that does not do a full page reload
HX-Push-Url - pushes a new url into the history stack
HX-Redirect - can be used to do a client-side redirect to a new location
HX-Refresh - if set to “true” the client-side will do a full refresh of the page
HX-Replace-Url - replaces the current URL in the location bar
HX-Reswap - allows you to specify how the response will be swapped. See hx-swap for possible values
HX-Retarget - a CSS selector that updates the target of the content update to a different element on the page
HX-Reselect - a CSS selector that allows you to choose which part of the response is used to be swapped in. Overrides an existing hx-select on the triggering element
HX-Trigger - allows you to trigger client-side events
HX-Trigger-After-Settle - allows you to trigger client-side events after the settle step
HX-Trigger-After-Swap - allows you to trigger client-side events after the swap step
*/
type HTMXResponseHeader string

const (
	HTMXLocation           HTMXResponseHeader = "HX-Location"
	HTMXPushUrl            HTMXResponseHeader = "HX-Push-Url"
	HTMXRedirect           HTMXResponseHeader = "HX-Redirect"
	HTMXRefresh            HTMXResponseHeader = "HX-Refresh"
	HTMXReplaceUrl         HTMXResponseHeader = "HX-Replace-Url"
	HTMXReswap             HTMXResponseHeader = "HX-Reswap"
	HTMXRetarget           HTMXResponseHeader = "HX-Retarget"
	HTMXReselect           HTMXResponseHeader = "HX-Reselect"
	HTMXTrigger            HTMXResponseHeader = "HX-Trigger"
	HTMXTriggerAfterSettle HTMXResponseHeader = "HX-Trigger-After-Settle"
	HTMXTriggerAfterSwap   HTMXResponseHeader = "HX-Trigger-After"
)

// This custom Render replaces Echo's echo.Context.Render() with templ's templ.Component.Render().
func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}

func HTMXAppendSuccessToast(c echo.Context, message string) {
	event := map[string]interface{}{
		"toast": map[string]interface{}{
			"message": message,
			"level":   "success",
		},
	}

	HTMXAppendEventsAfterSwap(c, event)
}

func HTMXAppendErrorToast(c echo.Context, message string) {
	event := map[string]interface{}{
		"toast": map[string]interface{}{
			"message": message,
			"level":   "error",
		},
	}

	HTMXAppendEventsAfterSwap(c, event)
}

func HTMXAppendToast(c echo.Context, message string, level string) {
	event := map[string]interface{}{
		"toast": map[string]interface{}{
			"message": message,
			"level":   level,
		},
	}

	HTMXAppendEventsAfterSwap(c, event)
}

func HTMXAppendPrelineRefresh(c echo.Context) {
	event := map[string]interface{}{
		"preline-refresh": "",
	}

	HTMXAppendEventsAfterSwap(c, event)
}

func HTMXCloseModal(c echo.Context) {
	event := map[string]interface{}{
		"close-modal": "",
	}

	HTMXAppendEventsAfterSwap(c, event)
}

func HTMXAppendTrigger(c echo.Context, trigger string) {
	// trigger, trigger, trigger
	hxTrigger := c.Response().Header().Get("HX-Trigger")
	if hxTrigger == "" {
		c.Response().Header().Set("HX-Trigger", trigger)
	} else {
		c.Response().Header().Set("HX-Trigger", hxTrigger+","+trigger)
	}
}

func HTMXAppendEventsAfterSwap(c echo.Context, events map[string]interface{}) {
	// Serialize the new events into JSON format
	eventTriggers := []string{}
	for key, value := range events {
		eventData, err := json.Marshal(value)
		if err != nil {
			// Handle JSON marshalling error (optional: log it or skip the event)
			continue
		}
		eventTriggers = append(eventTriggers, fmt.Sprintf(`"%s":%s`, key, eventData))
	}
	newTrigger := fmt.Sprintf("{%s}", strings.Join(eventTriggers, ","))

	// Append to existing "HX-Trigger-After-Swap" header
	existingTrigger := c.Response().Header().Get("HX-Trigger-After-Swap")
	if existingTrigger == "" {
		// No existing triggers, set the new one directly
		c.Response().Header().Set("HX-Trigger-After-Swap", newTrigger)
	} else {
		// Append the new trigger to the existing one
		mergedTrigger := existingTrigger[:len(existingTrigger)-1] + "," + newTrigger[1:]
		c.Response().Header().Set("HX-Trigger-After-Swap", mergedTrigger)
	}
}
