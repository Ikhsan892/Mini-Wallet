package service

import (
	"context"
	"log"
	"wallet/application/entity"
	"wallet/application/factory"
	"wallet/application/repositories/transaction"
	"wallet/application/repositories/wallet"
)

type GetTransactionService struct {
	ctx           context.Context
	walletTrxRepo transaction.TransactionRepository
	walletRepo    wallet.WalletRepository

	User entity.User
}

type GetTransactionServiceConfiguration func(m *GetTransactionService) error

func NewGetTransactionService(ctx context.Context, cfgs ...GetTransactionServiceConfiguration) *GetTransactionService {
	obj := &GetTransactionService{ctx: ctx}

	for _, cfg := range cfgs {
		err := cfg(obj)
		if err != nil {
			log.Fatal("Cannot hook with configuration")
		}
	}

	return obj
}

func WithWalletSqliteRepositoryForTransactionService(walletRepo wallet.WalletRepository) func(m *GetTransactionService) error {
	return func(m *GetTransactionService) error {
		m.walletRepo = walletRepo

		return nil
	}
}

func WithTransactionSqliteRepositoryForTransactionService(trxRepo transaction.TransactionRepository) func(m *GetTransactionService) error {
	return func(m *GetTransactionService) error {
		m.walletTrxRepo = trxRepo

		return nil
	}
}

func (t *GetTransactionService) Get() ([]factory.TransactionFactory, error) {
	w, err := t.walletRepo.GetWalletByOwner(t.ctx, t.User.Id)
	if err != nil {
		return []factory.TransactionFactory{}, err
	}

	trxs := t.walletTrxRepo.GetTransactionsByWallet(t.ctx, w.GetId())

	return t.toFactory(trxs), nil
}

func (t *GetTransactionService) toFactory(trxs []*entity.WalletTransaction) []factory.TransactionFactory {
	var result []factory.TransactionFactory

	if len(trxs) > 0 {
		for _, trx := range trxs {
			result = append(result, factory.NewTransactionFactory(trx))
		}
	}

	return result
}
