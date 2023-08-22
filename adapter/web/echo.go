package web

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.uber.org/zap"
	"net/http"
	"wallet"
	"wallet/adapter/web/api"
)

type EchoWebAdapter struct {
	ec  *echo.Echo
	app wallet.App
}

func NewEcho(app wallet.App) *EchoWebAdapter {
	e := echo.New()
	logger := app.ZapLogger()

	e.Debug = false

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info("request",
				zap.String("URI", v.URI),
				zap.Int("status", v.Status),
			)

			return nil
		},
	}))
	e.HTTPErrorHandler = HttpErrorHandler(app)
	e.Use(middleware.RequestID())
	e.Use(middleware.Gzip())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"localhost"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Content-Disposition"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderAuthorization, echo.HeaderContentType, "module", "Content-Range", "Accept-Language"},
	}))

	if app.Settings().App.Env == "development" {
		e.GET("/swagger/*", echoSwagger.WrapHandler)
	}

	api.InitRoutes(e, app)

	return &EchoWebAdapter{ec: e, app: app}
}

func (e *EchoWebAdapter) Init() error {

	s := http.Server{
		Addr:    fmt.Sprintf(":%s", e.app.Settings().App.AppWebserverPort),
		Handler: e.ec, // set Echo as handler
		//ReadTimeout: 30 * time.Second, // use custom timeouts
	}

	schema := "http"

	bold := color.New(color.Bold).Add(color.FgGreen)
	bold.Printf("> REST API Server started at: %s\n", color.CyanString("%s://localhost:%s", schema, e.app.Settings().App.AppWebserverPort))

	return s.ListenAndServe()
}
