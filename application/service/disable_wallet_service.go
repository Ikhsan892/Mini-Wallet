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
	ErrAccountAlreadyDisabled = errors.New("your wallet already disabled")
)

type DisableWalletService struct {
	ctx        context.Context
	walletRepo wallet.WalletRepository

	User entity.User
}

type DisableWalletServiceConfiguration func(m *DisableWalletService) error

func NewDisableWalletService(ctx context.Context, cfgs ...DisableWalletServiceConfiguration) *DisableWalletService {
	obj := &DisableWalletService{ctx: ctx}

	for _, cfg := range cfgs {
		err := cfg(obj)
		if err != nil {
			log.Fatal("Cannot hook with configuration")
		}
	}

	return obj
}

func WithWalletSqliteRepositoryForDisableWallet(walletRepo wallet.WalletRepository) func(m *DisableWalletService) error {
	return func(m *DisableWalletService) error {
		m.walletRepo = walletRepo

		return nil
	}
}

func (e *DisableWalletService) Update() (factory.WalletFactory, error) {
	w, err := e.walletRepo.GetWalletByOwner(e.ctx, e.User.Id)
	if err != nil {
		return factory.WalletFactory{}, err
	}

	if w.IsDisabled() {
		return factory.WalletFactory{}, ErrAccountAlreadyDisabled
	}

	w.DisableWallet()

	enabledWallet, err := e.walletRepo.UpdateWalletStatus(e.ctx, w)
	if err != nil {
		return factory.WalletFactory{}, err
	}

	return factory.NewWalletFactory(enabledWallet), nil
}
