package factory

import (
	"time"
	"wallet/application/entity"
)

type DepositFactory struct {
	Id          string    `json:"id"`
	DepositedBy string    `json:"deposited_by"`
	Status      string    `json:"status"`
	DepositedAt time.Time `json:"deposited_at"`
	Amount      float64   `json:"amount"`
	ReferenceId string    `json:"reference_id"`
}

func NewDepositFactory(trx *entity.WalletTransaction) DepositFactory {
	return DepositFactory{
		Id:          trx.GetId(),
		DepositedBy: trx.GetIssuedBy(),
		Status:      trx.GetStatus(),
		DepositedAt: trx.GetIssuedAt(),
		Amount:      trx.GetAmount(),
		ReferenceId: trx.GetRefId(),
	}
}
