package factory

import (
	"time"
	"wallet/application/entity"
)

type Withdrawn struct {
	Id          string    `json:"id"`
	WithdrawnBy string    `json:"withdrawn_by"`
	Status      string    `json:"status"`
	WithdrawnAt time.Time `json:"withdrawn_at"`
	Amount      float64   `json:"amount"`
	ReferenceId string    `json:"reference_id"`
}

func NewWithdrawn(trx *entity.WalletTransaction) Withdrawn {
	return Withdrawn{
		Id:          trx.GetId(),
		WithdrawnBy: trx.GetIssuedBy(),
		Status:      trx.GetStatus(),
		WithdrawnAt: trx.GetIssuedAt(),
		Amount:      trx.GetAmount(),
		ReferenceId: trx.GetRefId(),
	}
}
