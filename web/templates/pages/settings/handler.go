package settings

import (
	"log/slog"
	"net/http"

	"RouteHub.Service.Dashboard/ent"
	"RouteHub.Service.Dashboard/features/configuration"
	"RouteHub.Service.Dashboard/features/fileUpload"
	"RouteHub.Service.Dashboard/web/context"
	"RouteHub.Service.Dashboard/web/extensions"
	"RouteHub.Service.Dashboard/web/templates/pages/domain/components"
	"RouteHub.Service.Dashboard/web/templates/pages/partial"
	"github.com/labstack/echo/v4"
)

type Handlers struct {
	Ent               *ent.Client
	Logger            *slog.Logger
	ComponentHandlers *components.Handlers
}

func NewHandlers(logger *slog.Logger, ent *ent.Client) *Handlers {
	return &Handlers{
		Ent:               ent,
		Logger:            logger,
		ComponentHandlers: components.NewHandlers(logger, ent),
	}
}

func (h *Handlers) IndexHandler(c echo.Context) error {
	userInfo, _ := context.GetUserFromContext(c)
	currentConfig := configuration.Get()

	return extensions.Render(c, http.StatusOK, index(userInfo, currentConfig, nil))
}

func (h *Handlers) IndexPostHandler(c echo.Context) error {
	title := "Settings Form"
	currentConfig := configuration.Get()

	s3RequestPayload := configuration.S3ClientConfig{}

	if err := c.Bind(&s3RequestPayload); err != nil {
		feedback := partial.FormFeedbackFromErr(title, err)
		return extensions.Render(c, http.StatusOK, settingsPage(currentConfig, feedback))
	}

	_, err := s3RequestPayload.Validate()
	if err != nil {
		feedback := partial.FormFeedbackFromErr(title, err)
		return extensions.Render(c, http.StatusOK, settingsPage(currentConfig, feedback))
	}

	conf := configuration.Get()

	sc := conf.GetStaticConfig()
	sc.S3 = s3RequestPayload

	if err := conf.SetStaticConfig(sc); err != nil {
		feedback := partial.FormFeedbackFromErr(title, err)
		return extensions.Render(c, http.StatusOK, settingsPage(currentConfig, feedback))
	}

	_, err = fileUpload.GetS3ClientService()
	if err != nil {
		feedback := partial.FormFeedbackFromErr(title, err)
		return extensions.Render(c, http.StatusOK, settingsPage(currentConfig, feedback))
	}

	extensions.HTMXAppendSuccessToast(c, "Settings updated successfully")

	return extensions.Render(c, http.StatusOK, settingsPage(currentConfig, nil))
}
