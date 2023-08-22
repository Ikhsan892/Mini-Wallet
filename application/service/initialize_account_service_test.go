package service_test

import (
	"context"
	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/assert"
	"testing"
	"wallet/application/entity"
	"wallet/application/repositories/user"
	wallet2 "wallet/application/repositories/wallet"
	"wallet/application/service"
)

func TestInitializeAccountServiceSuccess(t *testing.T) {
	ctx := context.Background()
	id := ulid.Make().String()

	mockRepo := new(user.UserMockrepository)
	mockRepo.On("CreateUserAccount", ctx, id).Return(entity.User{Id: id}, nil)

	wallet := entity.NewWallet(id, id, "disabled", 0)
	mockWalletRepo := new(wallet2.WalletRepositoryMock)
	mockWalletRepo.On("CreateWallet", ctx, wallet).Return(wallet, nil)

	srv := service.NewInitializeAccountService(
		ctx,
		service.WithAccountSqliteRepository(mockRepo),
		service.WithWalletSqliteRepository(mockWalletRepo),
	)
	srv.CustId = id
	srv.WalletId = id

	data, err := srv.Submit()
	assert.Nil(t, err)

	assert.NotNil(t, data)

}
