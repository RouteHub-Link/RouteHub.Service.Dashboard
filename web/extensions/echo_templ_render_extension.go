package extensions

import (
	"fmt"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
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

func HTMXAppendToast(c echo.Context, message string) {
	trigger := fmt.Sprintf(`"toast":{"level" : "success", "message" : "%s"}`, message)

	if c.Response().Header().Get("HX-Trigger-After-Swap") == "" {
		c.Response().Header().Set("HX-Trigger-After-Swap", fmt.Sprintf(`{%s}`, trigger))
	} else {
		var hxTrigger = c.Response().Header().Get("HX-Trigger-After-Swap")
		hxTrigger = hxTrigger[:len(hxTrigger)-1] + "," + trigger + "}"
		c.Response().Header().Set("HX-Trigger-After-Swap", hxTrigger)
	}
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
