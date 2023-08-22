package api

import (
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"wallet"
	"wallet/application/repositories/user"
)

func InitRoutes(e *echo.Echo, app wallet.App) {
	ctx := context.Background()
	prefix := e.Group("/api")

	userRepo := user.NewUserSqliteRepository(app.DB())

	bindUserApi(prefix, ctx, app)
	prefix.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup:  "header:" + echo.HeaderAuthorization,
		AuthScheme: "Token",
		ErrorHandler: func(err error, c echo.Context) error {
			return NewUnauthorizedError(err.Error(), nil)
		},
		Validator: func(key string, c echo.Context) (bool, error) {
			if !userRepo.IsTokenValid(ctx, key) {
				return false, errors.New("Unauthorized")
			}

			us, _ := userRepo.GetUserAccountFromToken(ctx, key)
			c.Set("user", us)

			return true, nil
		},
	}))
	{
		bindWalletApi(prefix, ctx, app)
	}

}
