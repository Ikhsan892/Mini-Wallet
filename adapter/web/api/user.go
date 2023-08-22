package api

import (
	"context"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"wallet"
	"wallet/application/repositories/user"
	wallet2 "wallet/application/repositories/wallet"
	"wallet/application/service"
)

type UserController struct {
	ctx  context.Context
	conn *gorm.DB
}

// Initialize
// @Summary API untuk Inisialisasi akun baru
// @Tags    Menu
// @Accept  json
// @Produce json
// @Param   menus body     service.InitializeAccountService true "Data init"
// @Failure 400   {object} api.ApiError
// @Failure 401   {object} api.ApiError
// @Failure 404   {object} api.ApiError
// @Failure 500   {object} api.ApiError
// @Router  /v1/init [post]
func (m *UserController) InitializeAccount(c echo.Context) error {
	tx := m.conn.Begin()
	srv := service.NewInitializeAccountService(
		m.ctx,
		service.WithAccountSqliteRepository(user.NewUserSqliteRepository(tx)),
		service.WithWalletSqliteRepository(wallet2.NewWalletSqliteRepository(tx)),
	)

	if err := c.Bind(srv); err != nil {
		return NewForbiddenError("Error while binding", nil)
	}

	if err := srv.Validate(); err != nil {
		return NewValidationError(err)
	}

	defer tx.Rollback()

	token, err := srv.Submit()
	if err != nil {
		return NewBadRequestError(err.Error(), err)
	}

	tx.Commit()

	return NewApiResponse(200, map[string]interface{}{
		"token": token,
	}, c)
}

func bindUserApi(e *echo.Group, ctx context.Context, conn wallet.App) {
	userController := &UserController{
		ctx:  ctx,
		conn: conn.DB(),
	}

	prefix := e.Group("/v1/init")
	{
		prefix.POST("", userController.InitializeAccount)
	}

}
