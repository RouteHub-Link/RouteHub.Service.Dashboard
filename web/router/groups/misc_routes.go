package groups

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"RouteHub.Service.Dashboard/web/extensions"
	"RouteHub.Service.Dashboard/web/templates/pages/misc"
	"github.com/labstack/echo/v4"
)

func MapMiscRoutes(e *echo.Echo) {
	e.HTTPErrorHandler = customHTTPErrorHandler
	e.GET("/error", echo.WrapHandler(http.HandlerFunc(HandleError)))
	e.GET("/log", echo.WrapHandler(http.HandlerFunc(HandleLog)))
	e.GET("/data", echo.WrapHandler(http.HandlerFunc(HandleData)))

	e.GET("/404", handle404)

}

func handle404(c echo.Context) error {
	referer := c.QueryParam("referer")

	var referestr *string
	if referer != "" {
		if referer != c.Request().Host+c.Request().URL.Path {
			referestr = &referer
		}
	}

	return extensions.Render(c, http.StatusNotFound, misc.NotFound(referestr))
}

func customHTTPErrorHandler(err error, c echo.Context) {
	errorCode := err.(*echo.HTTPError).Code

	data := map[string]interface{}{}

	if errorCode != http.StatusOK {
		c.Logger().Print(err)
		c.Logger().Print(errorCode)

		if errorCode == http.StatusInternalServerError {
			data["code"] = "500"
			data["title"] = "Oops! Something went wrong."
			data["message"] = "Internal Server Error"

			//middlewares.ViewHandler(c, "error", "main", data)
		} else if errorCode == http.StatusNotFound {
			// make redirect url to referer path if available
			referer := c.Request().Referer()
			url := "/404"
			if referer != "" {
				url = fmt.Sprintf("%s?referer=%s", url, referer)
			}
			c.Redirect(http.StatusPermanentRedirect, url)
			return
			//middlewares.ViewHandler(c, "error", "main", data)
		} else if errorCode == http.StatusUnauthorized {
			data["code"] = "401"
			data["title"] = "Oops! Unauthorized."
			data["message"] = "You are not authorized to access this page."

			data["redirectURL"] = "/login"
			data["redirectTitle"] = "Login"

			//middlewares.ViewHandler(c, "error", "main", data)
		}

		dataAsJSON := json.NewEncoder(c.Response().Writer)
		dataAsJSON.SetIndent("", "  ")
		dataAsJSON.Encode(data)

		c.Render(errorCode, "error", dataAsJSON)
	}
}

func HandleLog(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "serial_id", "1")
	switch r.Method {
	case "GET":
		slog.InfoContext(ctx, "Handled GET request on /log",
			slog.Group("Request Info",
				slog.String("Method", "GET"),
				slog.String("Path", r.URL.Path),
			),
		)
		fmt.Fprintln(w, "Received a GET request at /log.")
	case "POST":
		slog.InfoContext(ctx, "Handled POST request on /log",
			slog.Group("Request Info",
				slog.String("Method", "POST"),
				slog.String("Path", r.URL.Path),
			),
		)
		fmt.Fprintln(w, "Received a POST request at /log.")
	default:
		http.Error(w, "Unsupported HTTP method", http.StatusMethodNotAllowed)
	}
}

func HandleData(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "request_id", fmt.Sprintf("%d", os.Getpid()))
	slog.InfoContext(ctx, "Data endpoint hit",
		slog.String("method", r.Method),
		slog.String("endpoint", "/data"),
	)
	fmt.Fprintln(w, "This is the data endpoint. Method used:", r.Method)
}

func HandleError(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "error_id", "error123")
	slog.ErrorContext(ctx, "Error endpoint accessed",
		slog.String("method", r.Method),
		slog.String("endpoint", "/error"),
	)
	http.Error(w, "You have reached the error endpoint", http.StatusInternalServerError)
}
