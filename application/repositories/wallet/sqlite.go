package wallet

import (
	"context"
	"gorm.io/gorm"
	"wallet/application/entity"
	"wallet/application/models"
)

type WalletSqliteRepository struct {
	conn *gorm.DB
}

func NewWalletSqliteRepository(conn *gorm.DB) *WalletSqliteRepository {
	return &WalletSqliteRepository{conn: conn}
}

func (a WalletSqliteRepository) GetWalletByOwner(ctx context.Context, owner string) (*entity.Wallet, error) {
	var data models.Wallet

	if err := a.conn.Model(&models.Wallet{}).WithContext(ctx).Where("owned_by = ?", owner).First(&data).Error; err != nil {
		return nil, err
	}

	return toAggregate(data), nil
}

func (a WalletSqliteRepository) UpdateWalletStatus(ctx context.Context, data *entity.Wallet) (*entity.Wallet, error) {
	if err := a.conn.Model(&models.Wallet{}).WithContext(ctx).Where("id = ?", data.GetId()).Updates(
		map[string]interface{}{
			"enabled_at": data.GetEnabledAt(),
			"status":     data.GetStatus(),
		},
	).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func toAggregate(w models.Wallet) *entity.Wallet {
	return entity.NewWallet(w.ID, w.OwnedBy, w.Status, w.Balance)
}

func (a WalletSqliteRepository) CreateWallet(ctx context.Context, data *entity.Wallet) (*entity.Wallet, error) {
	b, _ := data.GetBalance()

	wallet := models.Wallet{
		BaseModel: models.BaseModel{
			ID: data.GetId(),
		},
		OwnedBy:   data.OwnedBy(),
		Status:    data.GetStatus(),
		EnabledAt: data.GetEnabledAt(),
		Balance:   b,
	}

	if err := a.conn.WithContext(ctx).Model(&models.Wallet{}).Create(&wallet).Error; err != nil {
		return nil, err
	}

	return toAggregate(wallet), nil
}

func (a WalletSqliteRepository) UpdateBalance(ctx context.Context, amount float64, walletId string) error {
	if err := a.conn.WithContext(ctx).Model(&models.Wallet{}).Where("id = ?", walletId).Update("Balance", amount).Error; err != nil {
		return err
	}

	return nil
}
