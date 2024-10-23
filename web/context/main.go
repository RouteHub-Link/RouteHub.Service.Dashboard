package context

import (
	"log/slog"

	"RouteHub.Service.Dashboard/features"
	"github.com/labstack/echo/v4"
)

type ServerEchoContext struct {
	echo.Context
}

const (
	loggerKey         = "logger"
	clientServicesKey = "clientServices"
)

func ApplyMiddleware(e *echo.Echo) {
	lc := features.NewLoggerConfigurer(slog.LevelInfo)
	implementEchoLogger(e, lc, loggerKey)

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &ServerEchoContext{c}
			return next(cc)
		}
	})
}

func (cc *ServerEchoContext) GetLogger() *slog.Logger {
	return cc.Get(loggerKey).(*slog.Logger)
}
