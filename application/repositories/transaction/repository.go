package transaction

import (
	"context"
	"github.com/stretchr/testify/mock"
	"wallet/application/entity"
)

type TransactionRepository interface {
	GetTransactionsByWallet(ctx context.Context, walletId string) []*entity.WalletTransaction
	CreateTransaction(ctx context.Context, trx *entity.WalletTransaction) (*entity.WalletTransaction, error)
}

type MockTransactionRepository struct {
	mock.Mock
}

func (m *MockTransactionRepository) GetTransactionsByWallet(ctx context.Context, walletId string) []*entity.WalletTransaction {
	param := m.Called(ctx, walletId)

	return param.Get(0).([]*entity.WalletTransaction)
}

func (m *MockTransactionRepository) CreateTransaction(ctx context.Context, trx *entity.WalletTransaction) (*entity.WalletTransaction, error) {
	param := m.Called(ctx, trx)

	return param.Get(0).(*entity.WalletTransaction), param.Error(1)
}
