package groups

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func MapMiscRoutes(e *echo.Echo) {
	e.HTTPErrorHandler = customHTTPErrorHandler
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
			if referer != "" {

				// Check referer domain is same as current domain and wrote only path
				if referer != c.Request().Host+c.Request().URL.Path {

					data["redirectURL"] = referer
					data["redirectTitle"] = "Go Back"
					log.Printf("\n referer : %v", referer)
				}
			}

			data["code"] = "404"
			data["title"] = "Oops! Page not found."
			data["message"] = "We can't find the page you're looking for."

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
