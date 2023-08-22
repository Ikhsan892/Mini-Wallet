package service

import (
	"context"
	"github.com/go-playground/validator/v10"
	"log"
	"wallet/application/entity"
	"wallet/application/repositories/user"
	"wallet/application/repositories/wallet"
)

type InitializeAccountService struct {
	ctx        context.Context
	validate   *validator.Validate
	userRepo   user.UserRepository
	walletRepo wallet.WalletRepository

	CustId   string `form:"customer_xid" validate:"required"`
	WalletId string
}

type InitializeAccountServiceConfiguration func(m *InitializeAccountService) error

func NewInitializeAccountService(ctx context.Context, cfgs ...InitializeAccountServiceConfiguration) *InitializeAccountService {
	obj := &InitializeAccountService{ctx: ctx, validate: validator.New()}

	for _, cfg := range cfgs {
		err := cfg(obj)
		if err != nil {
			log.Fatal("Cannot hook with configuration")
		}
	}

	return obj
}

func WithAccountSqliteRepository(userRepo user.UserRepository) func(m *InitializeAccountService) error {
	return func(m *InitializeAccountService) error {

		m.userRepo = userRepo

		return nil
	}
}
func WithWalletSqliteRepository(walletRepo wallet.WalletRepository) func(m *InitializeAccountService) error {
	return func(m *InitializeAccountService) error {

		m.walletRepo = walletRepo

		return nil
	}
}

func (c *InitializeAccountService) Validate() error {
	err := c.validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}

func (c *InitializeAccountService) Submit() (string, error) {
	data, err := c.userRepo.CreateUserAccount(c.ctx, c.CustId)
	if err != nil {
		return "", err
	}

	_, err = c.walletRepo.CreateWallet(c.ctx, entity.NewWallet(c.WalletId, data.Id, "disabled", 0))
	if err != nil {
		return "", err
	}

	return data.Token, nil
}
