package web

import (
	"errors"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"wallet"
	"wallet/adapter/web/api"
)

func HttpErrorHandler(app wallet.App) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		report, ok := err.(*echo.HTTPError)

		var errorApp *api.ApiError
		if !ok {
			if errors.As(err, &errorApp) {
				report = echo.NewHTTPError(errorApp.Code, errorApp.Data)
			} else {
				if app.Settings().App.Env == "production" {
					msg := "Oops something went wrong, Contact Developer for the issue"
					report = echo.NewHTTPError(http.StatusInternalServerError, msg)
				} else {
					report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
				}
				app.ZapLogger().Error(err.Error())
			}
		}

		err2 := c.JSON(report.Code, errorApp)
		if err2 != nil {
			log.Fatal("Cannot returna anything")
		}
	}
}
