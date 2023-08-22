package service

import (
	"context"
	"log"
	"wallet/application/entity"
	"wallet/application/factory"
	"wallet/application/repositories/transaction"
	"wallet/application/repositories/wallet"
	"wallet/application/value_objects"
)

type AddAmountService struct {
	ctx           context.Context
	walletRepo    wallet.WalletRepository
	walletTrxRepo transaction.TransactionRepository

	Amount float64 `form:"amount"`
	RefId  string  `form:"reference_id"`
	User   entity.User
}

type AddAmountServiceConfiguration func(m *AddAmountService) error

func NewAddAmountService(ctx context.Context, cfgs ...AddAmountServiceConfiguration) *AddAmountService {
	obj := &AddAmountService{ctx: ctx}

	for _, cfg := range cfgs {
		err := cfg(obj)
		if err != nil {
			log.Fatal("Cannot hook with configuration")
		}
	}

	return obj
}

func WithWalletSqliteRepositoryForAddAmount(walletRepo wallet.WalletRepository) func(m *AddAmountService) error {
	return func(m *AddAmountService) error {
		m.walletRepo = walletRepo

		return nil
	}
}

func WithTransactionSqliteRepositoryForAddAmount(trxRepo transaction.TransactionRepository) func(m *AddAmountService) error {
	return func(m *AddAmountService) error {
		m.walletTrxRepo = trxRepo

		return nil
	}
}

func (a *AddAmountService) Submit() (factory.DepositFactory, error) {
	w, err := a.walletRepo.GetWalletByOwner(a.ctx, a.User.Id)
	if err != nil {
		return factory.DepositFactory{}, err
	}

	err = w.AddBalance(value_objects.NewBalance(a.Amount))
	if err != nil {
		return factory.DepositFactory{}, err
	}

	trx, err := entity.NewWalletTransaction(w.GetId(), "DEPOSIT", a.User.Id, "success", a.Amount)
	if err != nil {
		return factory.DepositFactory{}, err
	}

	createdTransaction, err := a.walletTrxRepo.CreateTransaction(a.ctx, trx)
	if err != nil {
		return factory.DepositFactory{}, err
	}

	b, _ := w.GetBalance()
	err = a.walletRepo.UpdateBalance(a.ctx, b, w.GetId())
	if err != nil {
		return factory.DepositFactory{}, err
	}

	return factory.NewDepositFactory(createdTransaction), nil

}
