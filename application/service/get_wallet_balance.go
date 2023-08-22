package service

import (
	"context"
	"log"
	"wallet/application/entity"
	"wallet/application/factory"
	"wallet/application/repositories/wallet"
)

type GetWalletBalanceService struct {
	ctx        context.Context
	walletRepo wallet.WalletRepository

	User entity.User
}

type GetWalletBalanceServiceConfiguration func(m *GetWalletBalanceService) error

func NewGetWalletBalanceService(ctx context.Context, cfgs ...GetWalletBalanceServiceConfiguration) *GetWalletBalanceService {
	obj := &GetWalletBalanceService{ctx: ctx}

	for _, cfg := range cfgs {
		err := cfg(obj)
		if err != nil {
			log.Fatal("Cannot hook with configuration")
		}
	}

	return obj
}

func WithWalletSqliteRepositoryForGetWalletBalance(walletRepo wallet.WalletRepository) func(m *GetWalletBalanceService) error {
	return func(m *GetWalletBalanceService) error {
		m.walletRepo = walletRepo

		return nil
	}
}

func (g *GetWalletBalanceService) Get() (factory.WalletFactory, error) {
	w, err := g.walletRepo.GetWalletByOwner(g.ctx, g.User.Id)
	if err != nil {
		return factory.WalletFactory{}, err
	}
	_, err = w.GetBalance()
	if err != nil {
		return factory.WalletFactory{}, err
	}

	return factory.NewWalletFactory(w), nil
}
