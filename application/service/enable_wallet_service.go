package service

import (
	"context"
	"errors"
	"log"
	"wallet/application/entity"
	"wallet/application/factory"
	"wallet/application/repositories/wallet"
)

var (
	ErrAccountAlreadyEnabled = errors.New("your wallet already enabled")
)

type EnableWalletService struct {
	ctx        context.Context
	walletRepo wallet.WalletRepository

	User entity.User
}

type EnableWalletServiceConfiguration func(m *EnableWalletService) error

func NewEnableWalletService(ctx context.Context, cfgs ...EnableWalletServiceConfiguration) *EnableWalletService {
	obj := &EnableWalletService{ctx: ctx}

	for _, cfg := range cfgs {
		err := cfg(obj)
		if err != nil {
			log.Fatal("Cannot hook with configuration")
		}
	}

	return obj
}

func WithWalletSqliteRepositoryForEnableWallet(walletRepo wallet.WalletRepository) func(m *EnableWalletService) error {
	return func(m *EnableWalletService) error {
		m.walletRepo = walletRepo

		return nil
	}
}

func (e *EnableWalletService) Update() (factory.WalletFactory, error) {
	w, err := e.walletRepo.GetWalletByOwner(e.ctx, e.User.Id)
	if err != nil {
		return factory.WalletFactory{}, err
	}

	if w.IsEnabled() {
		return factory.WalletFactory{}, ErrAccountAlreadyEnabled
	}

	w.EnabledWallet()

	enabledWallet, err := e.walletRepo.UpdateWalletStatus(e.ctx, w)
	if err != nil {
		return factory.WalletFactory{}, err
	}

	return factory.NewWalletFactory(enabledWallet), nil
}
