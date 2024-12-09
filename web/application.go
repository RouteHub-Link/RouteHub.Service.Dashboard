package web

import (
	"context"
	"errors"
	"log"
	"log/slog"
	"strings"
	"sync"

	"RouteHub.Service.Dashboard/ent"
	"RouteHub.Service.Dashboard/ent/migrate"
	"RouteHub.Service.Dashboard/features"
	"RouteHub.Service.Dashboard/features/auth"
	"RouteHub.Service.Dashboard/features/configuration"
	"RouteHub.Service.Dashboard/web/extensions"
	"RouteHub.Service.Dashboard/web/middlewares"
	"RouteHub.Service.Dashboard/web/router"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	slogecho "github.com/samber/slog-echo"
	otel_middleware "go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
)

var (
	onceApplication sync.Once
	application     *Application

	ErrApplicationNotInitialized = errors.New("Application not initialized")
)

// Application structure to hold all components
type Application struct {
	Echo          *echo.Echo
	config        *configuration.Config
	oauthInstance *auth.OAuthInstance
	Authorizer    *extensions.Authorizer
	Logger        *slog.Logger
	Ent           *ent.Client
}

func GetApplication() (*Application, error) {
	if application == nil {
		return nil, ErrApplicationNotInitialized
	}

	return application, nil
}

func NewApplication(config *configuration.Config, logger *slog.Logger) (*Application, error) {
	onceApplication.Do(func() {
		// Initialize Echo
		e := echo.New()
		e.Validator = &extensions.CustomValidator{Validator: validator.New()}
		e.Server.Addr = strings.Join([]string{config.Server.Host, config.Server.Port}, ":")
		logger.Info("Server Configured", "address", e.Server.Addr)

		// Initialize Authorizer
		oauthInstance, err := auth.NewOAuthInstance(auth.WithOAuthConfig(config.OAuth), auth.WithLogger(logger))
		if err != nil {
			logger.Error("Failed to create OAuthInstance", "error", err)
			return
		}

		authorizer, err := extensions.NewAuthorizer(oauthInstance)
		if err != nil {
			logger.Error("Failed to create Authorizer", "error", err)
			return
		}

		app := &Application{
			Echo:          e,
			Authorizer:    authorizer,
			Logger:        logger,
			config:        config,
			oauthInstance: oauthInstance,
		}

		if err := app.configureDatabase(); err != nil {
			logger.Error("Failed to configure database", "error", err)
			return
		}

		if _, err := features.NewCacheService(config.CacheDB, logger); err != nil {
			logger.Error("Failed to configure cache", "error", err)
			return
		}

		app.configureMiddleware()

		application = app
	})

	return application, nil
}

func (app *Application) configureMiddleware() {
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

	router.ConfigureRoutes(e, app.config, app.Logger, app.Authorizer, app.Ent)

	// TODO When this middleware is added, the application gets an error when trying to access the /auth/logout
	app.Echo.Use(middlewares.OAuthGuard(app.Authorizer, app.config.OAuth, app.Logger))

	app.Echo.Use(middlewares.PersonMiddleware(app.Authorizer, app.Logger, app.Ent))

	app.Logger.Info("OTEL Status", "enabled", app.config.OTEL.IsEnabled())

	if app.config.OTEL.IsEnabled() {
		app.configureOTEL()
	}
}

func (app *Application) configureOTEL() {
	otelConfig := app.config.OTEL

	appName := otelConfig.GetServiceName()
	openTelemetryOptions := []features.OpenTelemetryOption{
		features.WithInsecure(otelConfig.IsInsecureAsString()),
		features.WithServiceName(appName),
		features.WithCollectorURL(otelConfig.GetCollectorEndpoint()),
		features.WithLogger(app.Logger),
		features.WithHeaders(otelConfig.GetHeaders()),
	}

	ots := features.NewOpenTelemetryService(openTelemetryOptions...)
	ots.InitTracer()

	app.Echo.Use(otel_middleware.Middleware(appName))
}

func (app *Application) configureDatabase() error {
	ctx := context.Background()
	dbConfig := app.config.Database
	client, err := ent.Open("postgres", dbConfig.GetConnectionString())
	if err != nil {
		return err
	}

	if err := client.Schema.Create(ctx, migrate.WithGlobalUniqueID(false)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	app.Ent = client

	return nil
}
