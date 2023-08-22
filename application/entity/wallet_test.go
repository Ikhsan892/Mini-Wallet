package entity_test

import (
	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/assert"
	"testing"
	"wallet/application/entity"
	"wallet/application/value_objects"
)

func TestAddBalance(t *testing.T) {
	ownedBy := ulid.Make().String()
	testTables := []struct {
		name            string
		balance         float64
		currBalance     float64
		wantErr         bool
		status          string
		err             error
		expectedBalance float64
	}{
		{
			name:            "Success add balance",
			balance:         6000.00,
			currBalance:     2000.00,
			wantErr:         false,
			status:          "enabled",
			err:             nil,
			expectedBalance: 8000.00,
		},
		{
			name:            "Error add balance because account is still disable",
			balance:         6000.00,
			currBalance:     0,
			wantErr:         true,
			status:          "disabled",
			err:             entity.ErrWalletDisable,
			expectedBalance: 0,
		},
	}

	for _, table := range testTables {
		t.Run(table.name, func(t *testing.T) {
			wallet := entity.NewWallet("", ownedBy, table.status, table.currBalance)

			err := wallet.AddBalance(value_objects.NewBalance(table.balance))

			if table.wantErr {
				assert.ErrorIs(t, err, table.err)
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				b, _ := wallet.GetBalance()
				assert.Equal(t, b, table.expectedBalance)
			}
		})

	}
}

func TestGetBalance(t *testing.T) {
	ownedBy := ulid.Make().String()
	w := entity.NewWallet("", ownedBy, "disabled", 0)

	_, err := w.GetBalance()
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, entity.ErrWalletDisable)
}

func TestDeductBalance(t *testing.T) {
	ownedBy := ulid.Make().String()
	testTables := []struct {
		name            string
		balance         float64
		currBalance     float64
		wantErr         bool
		status          string
		err             error
		expectedBalance float64
	}{
		{
			name:            "Success deduct balance",
			balance:         6000.00,
			currBalance:     7000.00,
			wantErr:         false,
			status:          "enabled",
			err:             nil,
			expectedBalance: 1000.00,
		},
		{
			name:            "Error deduct balance because account is still disable",
			balance:         6000.00,
			currBalance:     0,
			wantErr:         true,
			status:          "disabled",
			err:             entity.ErrWalletDisable,
			expectedBalance: 0,
		},
		{
			name:            "Error deduct balance because amount is insufficient",
			balance:         6000.00,
			currBalance:     0,
			wantErr:         true,
			status:          "enabled",
			err:             entity.ErrBalanceInsufficient,
			expectedBalance: 0,
		},
	}

	for _, table := range testTables {
		t.Run(table.name, func(t *testing.T) {
			wallet := entity.NewWallet("", ownedBy, table.status, table.currBalance)

			err := wallet.DeductBalance(value_objects.NewBalance(table.balance))

			if table.wantErr {
				assert.ErrorIs(t, err, table.err)
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				b, _ := wallet.GetBalance()
				assert.Equal(t, b, table.expectedBalance)
			}
		})

	}
}
