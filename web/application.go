package web

import (
	"log/slog"
	"strings"

	"RouteHub.Service.Dashboard/features"
	"RouteHub.Service.Dashboard/features/auth"
	"RouteHub.Service.Dashboard/features/configuration"
	"RouteHub.Service.Dashboard/web/extensions"
	"RouteHub.Service.Dashboard/web/middlewares"
	"RouteHub.Service.Dashboard/web/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	slogecho "github.com/samber/slog-echo"
	otel_middleware "go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
)

// Application structure to hold all components
type Application struct {
	Echo          *echo.Echo
	config        *configuration.Config
	oauthInstance *auth.OAuthInstance
	Authorizer    *extensions.Authorizer
	Logger        *slog.Logger
}

func NewApplication(config *configuration.Config, logger *slog.Logger) (*Application, error) {
	// Initialize Echo
	e := echo.New()
	e.Server.Addr = strings.Join([]string{config.Server.Host, config.Server.Port}, ":")
	logger.Info("Server Configured", "address", e.Server.Addr)

	// Initialize Authorizer
	oauthInstance, err := auth.NewOAuthInstance(auth.WithOAuthConfig(config.OAuth), auth.WithLogger(logger))
	if err != nil {
		return nil, err
	}

	authorizer, err := extensions.NewAuthorizer(oauthInstance)
	if err != nil {
		return nil, err
	}

	app := &Application{
		Echo:          e,
		Authorizer:    authorizer,
		Logger:        logger,
		config:        config,
		oauthInstance: oauthInstance,
	}

	app.ConfigureMiddleware()

	return app, nil
}

func (app *Application) ConfigureMiddleware() {
	loggingConfig := slogecho.Config{
		WithSpanID:  true,
		WithTraceID: true,
	}

	e := app.Echo

	e.Use(slogecho.NewWithConfig(app.Logger, loggingConfig))

	e.Use(middleware.Recover())
	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		XSSProtection:         "1; mode=block",
		ContentTypeNosniff:    "nosniff",
		XFrameOptions:         "DENY",
		HSTSMaxAge:            31536000,
		HSTSExcludeSubdomains: true,
	}))

	// TODO When this middleware is added, the application gets an error when trying to access the /auth/logout
	app.Echo.Use(middlewares.OAuthGuard(app.Authorizer, app.config.OAuth, app.Logger))

	router.ConfigureRoutes(e, app.config, app.Authorizer)

	app.Logger.Info("OTEL Status", "enabled", app.config.OTEL.IsEnabled())

	if app.config.OTEL.IsEnabled() {
		app.ConfigureOTEL()
	}
}

func (app *Application) ConfigureOTEL() {
	otelConfig := app.config.OTEL

	appName := otelConfig.GetServiceName()
	openTelemetryOptions := []features.OpenTelemetryOption{
		features.WithInsecure(otelConfig.IsInsecureAsString()),
		features.WithServiceName(appName),
		features.WithCollectorURL(otelConfig.GetCollectorEndpoint()),
		features.WithLogger(app.Logger),
	}

	ots := features.NewOpenTelemetryService(openTelemetryOptions...)
	ots.InitTracer()

	app.Echo.Use(otel_middleware.Middleware(appName))
}
