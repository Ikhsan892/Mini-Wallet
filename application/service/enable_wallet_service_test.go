package service_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"wallet/application/entity"
	wallet2 "wallet/application/repositories/wallet"
	"wallet/application/service"
)

// TODO(ikhsan) : Tambahin test error pliss!!!
func TestEnableWallet(t *testing.T) {
	ctx := context.Background()
	user := entity.User{
		Id:    "123",
		Token: "123",
	}

	w := entity.NewWallet("", user.Id, "disabled", 0)

	mockWallet := new(wallet2.WalletRepositoryMock)
	mockWallet.On("GetWalletByOwner", ctx, user.Id).Return(w, nil)
	mockWallet.On("EnableWallet", ctx, w).Return(w, nil)
	srv := service.NewEnableWalletService(
		ctx,
		service.WithWalletSqliteRepositoryForEnableWallet(mockWallet),
	)

	srv.User = user

	wallet, err := srv.Update()
	assert.Nil(t, err)
	assert.Equal(t, "enabled", wallet.Status)

}
