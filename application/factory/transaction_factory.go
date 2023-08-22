package factory

import (
	"time"
	"wallet/application/entity"
)

type TransactionFactory struct {
	WalletId    string    `json:"wallet_id"`
	Type        string    `json:"type" gorm:"not null"`
	Amount      float64   `json:"amount"`
	IssuedBy    string    `json:"issued_by"`
	ReferenceId string    `json:"reference_id"`
	Status      string    `json:"status"`
	IssuedAt    time.Time `json:"issued_at"`
}

func NewTransactionFactory(trx *entity.WalletTransaction) TransactionFactory {
	return TransactionFactory{
		WalletId:    trx.GetWalletId(),
		Type:        trx.GetType(),
		Amount:      trx.GetAmount(),
		IssuedBy:    trx.GetIssuedBy(),
		ReferenceId: trx.GetRefId(),
		Status:      trx.GetStatus(),
		IssuedAt:    trx.GetIssuedAt(),
	}
}
