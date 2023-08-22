package factory

import (
	"time"
	"wallet/application/entity"
)

type WalletFactory struct {
	Id        string     `json:"id"`
	OwnedBy   string     `json:"owned_by"`
	Status    string     `json:"status"`
	EnabledAt *time.Time `json:"enabled_at"`
	Balance   float64    `json:"balance"`
}

func NewWalletFactory(w *entity.Wallet) WalletFactory {
	b, _ := w.GetBalance()

	return WalletFactory{
		Id:        w.GetId(),
		OwnedBy:   w.OwnedBy(),
		Status:    w.GetStatus(),
		EnabledAt: w.GetEnabledAt(),
		Balance:   b,
	}
}
