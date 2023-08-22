package entity_test

import (
	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/assert"
	"testing"
	"wallet/application/entity"
)

func TestNewTransaction(t *testing.T) {
	walletId := ulid.Make().String()

	trx, err := entity.NewWalletTransaction(walletId, "WITHDRAWALS", "123", "success", 6000)

	assert.Equal(t, trx.GetWalletId(), walletId)
	assert.Nil(t, err)

}

func TestNewTransactionUnknownType(t *testing.T) {
	walletId := ulid.Make().String()

	_, err := entity.NewWalletTransaction(walletId, "feawafweaf", "123", "failed", 6000)

	assert.ErrorIs(t, err, entity.ErrTypeIsUnknown)
	assert.NotNil(t, err)

}
