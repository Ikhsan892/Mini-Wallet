package wallet

import (
	"context"
	"github.com/stretchr/testify/mock"
	"wallet/application/entity"
)

type WalletRepository interface {
	CreateWallet(ctx context.Context, data *entity.Wallet) (*entity.Wallet, error)
	GetWalletByOwner(ctx context.Context, owner string) (*entity.Wallet, error)
	UpdateWalletStatus(ctx context.Context, data *entity.Wallet) (*entity.Wallet, error)
	UpdateBalance(ctx context.Context, amount float64, walletId string) error
}

type WalletRepositoryMock struct {
	mock.Mock
}

func (a *WalletRepositoryMock) CreateWallet(ctx context.Context, data *entity.Wallet) (*entity.Wallet, error) {
	param := a.Called(ctx, data)

	return param.Get(0).(*entity.Wallet), param.Error(1)
}

func (a *WalletRepositoryMock) GetWalletByOwner(ctx context.Context, owner string) (*entity.Wallet, error) {
	param := a.Called(ctx, owner)

	return param.Get(0).(*entity.Wallet), param.Error(1)
}

func (a *WalletRepositoryMock) UpdateWalletStatus(ctx context.Context, data *entity.Wallet) (*entity.Wallet, error) {
	param := a.Called(ctx, data)

	return param.Get(0).(*entity.Wallet), param.Error(1)
}

func (a *WalletRepositoryMock) UpdateBalance(ctx context.Context, amount float64, walletId string) error {
	param := a.Called(ctx, amount)

	return param.Error(0)
}
