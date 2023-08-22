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

type DeductAmountService struct {
	ctx           context.Context
	walletRepo    wallet.WalletRepository
	walletTrxRepo transaction.TransactionRepository

	Amount float64 `form:"amount"`
	RefId  string  `form:"reference_id"`
	User   entity.User
}

type DeductAmountServiceConfiguration func(m *DeductAmountService) error

func NewDeductAmountService(ctx context.Context, cfgs ...DeductAmountServiceConfiguration) *DeductAmountService {
	obj := &DeductAmountService{ctx: ctx}

	for _, cfg := range cfgs {
		err := cfg(obj)
		if err != nil {
			log.Fatal("Cannot hook with configuration")
		}
	}

	return obj
}

func WithWalletSqliteRepositoryForDeductAmount(walletRepo wallet.WalletRepository) func(m *DeductAmountService) error {
	return func(m *DeductAmountService) error {
		m.walletRepo = walletRepo

		return nil
	}
}

func WithTransactionSqliteRepositoryForDeductAmount(trxRepo transaction.TransactionRepository) func(m *DeductAmountService) error {
	return func(m *DeductAmountService) error {
		m.walletTrxRepo = trxRepo

		return nil
	}
}

func (a *DeductAmountService) Submit() (factory.Withdrawn, error) {
	w, err := a.walletRepo.GetWalletByOwner(a.ctx, a.User.Id)
	if err != nil {
		return factory.Withdrawn{}, err
	}

	err = w.DeductBalance(value_objects.NewBalance(a.Amount))
	if err != nil {
		return factory.Withdrawn{}, err
	}

	trx, err := entity.NewWalletTransaction(w.GetId(), "WITHDRAWALS", a.User.Id, "success", a.Amount)
	if err != nil {
		return factory.Withdrawn{}, err
	}

	createdTransaction, err := a.walletTrxRepo.CreateTransaction(a.ctx, trx)
	if err != nil {
		return factory.Withdrawn{}, err
	}

	b, _ := w.GetBalance()
	err = a.walletRepo.UpdateBalance(a.ctx, b, w.GetId())
	if err != nil {
		return factory.Withdrawn{}, err
	}

	return factory.NewWithdrawn(createdTransaction), nil

}
