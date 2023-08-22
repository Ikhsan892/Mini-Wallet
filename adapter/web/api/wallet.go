package api

import (
	"context"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"wallet"
	"wallet/application/entity"
	"wallet/application/repositories/transaction"
	wallet2 "wallet/application/repositories/wallet"
	"wallet/application/service"
)

type WalletController struct {
	ctx  context.Context
	conn *gorm.DB
}

func (w *WalletController) EnableMyWallet(c echo.Context) error {
	srv := service.NewEnableWalletService(
		w.ctx,
		service.WithWalletSqliteRepositoryForEnableWallet(
			wallet2.NewWalletSqliteRepository(w.conn),
		),
	)

	srv.User = c.Get("user").(entity.User)

	data, err := srv.Update()
	if err != nil {
		return NewBadRequestError(err.Error(), nil)
	}

	return NewApiResponse(200, map[string]interface{}{
		"wallet": data,
	}, c)
}

func (w *WalletController) DisableMyWallet(c echo.Context) error {
	srv := service.NewDisableWalletService(
		w.ctx,
		service.WithWalletSqliteRepositoryForDisableWallet(
			wallet2.NewWalletSqliteRepository(w.conn),
		),
	)

	srv.User = c.Get("user").(entity.User)

	data, err := srv.Update()
	if err != nil {
		return NewBadRequestError(err.Error(), nil)
	}

	return NewApiResponse(200, map[string]interface{}{
		"wallet": data,
	}, c)
}

func (w *WalletController) GetBalance(c echo.Context) error {
	srv := service.NewGetWalletBalanceService(
		w.ctx,
		service.WithWalletSqliteRepositoryForGetWalletBalance(
			wallet2.NewWalletSqliteRepository(w.conn),
		),
	)

	srv.User = c.Get("user").(entity.User)

	data, err := srv.Get()
	if err != nil {
		return NewBadRequestError(err.Error(), nil)
	}

	return NewApiResponse(200, map[string]interface{}{
		"wallet": data,
	}, c)

}

func (w *WalletController) AddBalance(c echo.Context) error {
	srv := service.NewAddAmountService(
		w.ctx,
		service.WithTransactionSqliteRepositoryForAddAmount(
			transaction.NewTransactionSqliteRepository(w.conn),
		),
		service.WithWalletSqliteRepositoryForAddAmount(
			wallet2.NewWalletSqliteRepository(w.conn),
		),
	)

	if err := c.Bind(srv); err != nil {
		return NewBadRequestError("cannot binding", err)
	}

	srv.User = c.Get("user").(entity.User)

	data, err := srv.Submit()
	if err != nil {
		return NewBadRequestError(err.Error(), nil)
	}

	return NewApiResponse(200, map[string]interface{}{
		"deposit": data,
	}, c)

}

func (w *WalletController) DeductBalance(c echo.Context) error {
	srv := service.NewDeductAmountService(
		w.ctx,
		service.WithTransactionSqliteRepositoryForDeductAmount(
			transaction.NewTransactionSqliteRepository(w.conn),
		),
		service.WithWalletSqliteRepositoryForDeductAmount(
			wallet2.NewWalletSqliteRepository(w.conn),
		),
	)

	if err := c.Bind(srv); err != nil {
		return NewBadRequestError("cannot binding", err)
	}

	srv.User = c.Get("user").(entity.User)

	data, err := srv.Submit()
	if err != nil {
		return NewBadRequestError(err.Error(), nil)
	}

	return NewApiResponse(200, map[string]interface{}{
		"deposit": data,
	}, c)

}

func (w *WalletController) GetTransactions(c echo.Context) error {
	srv := service.NewGetTransactionService(
		w.ctx,
		service.WithTransactionSqliteRepositoryForTransactionService(
			transaction.NewTransactionSqliteRepository(w.conn),
		),
		service.WithWalletSqliteRepositoryForTransactionService(
			wallet2.NewWalletSqliteRepository(w.conn),
		),
	)

	srv.User = c.Get("user").(entity.User)

	data, err := srv.Get()
	if err != nil {
		return NewBadRequestError(err.Error(), nil)
	}

	return NewApiResponse(200, map[string]interface{}{
		"transactions": data,
	}, c)

}

func bindWalletApi(e *echo.Group, ctx context.Context, conn wallet.App) {
	walletController := &WalletController{
		ctx:  ctx,
		conn: conn.DB(),
	}

	prefix := e.Group("/v1/wallet")
	{
		prefix.POST("", walletController.EnableMyWallet)
		prefix.GET("", walletController.GetBalance)
		prefix.PATCH("", walletController.DisableMyWallet)
		prefix.POST("/deposits", walletController.AddBalance)
		prefix.POST("/withdrawals", walletController.DeductBalance)
		prefix.GET("/transactions", walletController.GetTransactions)
	}

}
