package entity

import (
	"errors"
	"github.com/oklog/ulid/v2"
	"time"
)

var (
	ErrTypeIsUnknown = errors.New("transaction type is unknown")
)

type WalletTransaction struct {
	id        string
	walletId  string
	type_     string
	amount    float64
	status    string
	issued_at time.Time
	issued_by string
	ref_id    string
}

func NewWalletTransaction(walletId, type_, issued_by, status string, amount float64) (*WalletTransaction, error) {
	if type_ != "WITHDRAWALS" && type_ != "DEPOSIT" {
		return nil, ErrTypeIsUnknown
	}

	return &WalletTransaction{
		id:        ulid.Make().String(),
		walletId:  walletId,
		type_:     type_,
		amount:    amount,
		issued_at: time.Now(),
		issued_by: issued_by,
		status:    status,
		ref_id:    ulid.Make().String(),
	}, nil
}

func (trx WalletTransaction) GetIssuedAt() time.Time {
	return trx.issued_at
}

func (trx WalletTransaction) GetRefId() string {
	return trx.ref_id
}

func (trx WalletTransaction) GetStatus() string {
	return trx.status
}

func (trx WalletTransaction) GetIssuedBy() string {
	return trx.issued_by
}

func (trx WalletTransaction) GetId() string {
	return trx.id
}
func (trx WalletTransaction) GetWalletId() string {
	return trx.walletId
}
func (trx WalletTransaction) GetType() string {
	return trx.type_
}
func (trx WalletTransaction) GetAmount() float64 {
	return trx.amount
}
