package transaction

import (
	"context"
	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
	"wallet/application/entity"
	"wallet/application/models"
)

type TransactionSqliteRepository struct {
	conn *gorm.DB
}

func NewTransactionSqliteRepository(conn *gorm.DB) *TransactionSqliteRepository {
	return &TransactionSqliteRepository{conn}
}

func toAggregate(m models.Transaction) *entity.WalletTransaction {
	w, _ := entity.NewWalletTransaction(m.WalletId, m.Type, m.IssuedBy, m.Status, m.Amount)
	return w
}

func (t TransactionSqliteRepository) GetTransactionsByWallet(ctx context.Context, walletId string) []*entity.WalletTransaction {
	var result []models.Transaction
	var domains []*entity.WalletTransaction

	t.conn.Model(&models.Transaction{}).WithContext(ctx).Where("wallet_id = ?", walletId).Find(&result)

	if len(result) > 0 {
		for _, transaction := range result {
			domains = append(domains, toAggregate(transaction))
		}
	}

	return domains
}

func (t TransactionSqliteRepository) CreateTransaction(ctx context.Context, trx *entity.WalletTransaction) (*entity.WalletTransaction, error) {
	result := models.Transaction{
		BaseModel: models.BaseModel{
			ID: ulid.Make().String(),
		},
		WalletId:    trx.GetWalletId(),
		Type:        trx.GetType(),
		Amount:      trx.GetAmount(),
		IssuedBy:    trx.GetIssuedBy(),
		ReferenceId: trx.GetRefId(),
		Status:      trx.GetStatus(),
	}

	if err := t.conn.Model(&models.Transaction{}).WithContext(ctx).Create(&result).Error; err != nil {
		return nil, err
	}

	return toAggregate(result), nil
}
