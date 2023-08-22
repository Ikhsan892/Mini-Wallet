package entity

import (
	"errors"
	"github.com/oklog/ulid/v2"
	"time"
	"wallet/application/value_objects"
)

var (
	ErrBalanceInsufficient = errors.New("balance insufficient")
	ErrWalletDisable       = errors.New("Wallet Disabled")
)

type Wallet struct {
	id         string
	balance    value_objects.Balance
	status     string
	enabled_at *time.Time
	owned_by   string
}

func NewWallet(id, owned_by, status string, amount float64) *Wallet {
	ids := id
	if id == "" {
		ids = ulid.Make().String()
	}

	return &Wallet{
		id:       ids,
		balance:  value_objects.NewBalance(amount),
		status:   status,
		owned_by: owned_by,
	}
}

func (w *Wallet) EnabledWallet() {
	now := time.Now()

	w.enabled_at = &now
	w.status = "enabled"
}

func (w *Wallet) DisableWallet() {

	w.enabled_at = nil
	w.status = "disabled"
}

func (w *Wallet) SetBalance(balance value_objects.Balance) {
	w.balance = balance
}

func (w *Wallet) AddBalance(balance value_objects.Balance) error {
	if w.IsDisabled() {
		return ErrWalletDisable
	}

	w.balance = value_objects.NewBalance(w.balance.GetAmount() + balance.GetAmount())

	return nil
}

func (w Wallet) IsEnabled() bool {
	return w.GetStatus() == "enabled"
}

func (w *Wallet) DeductBalance(balance value_objects.Balance) error {
	if w.IsDisabled() {
		return ErrWalletDisable
	}

	if w.balance.GetAmount() < balance.GetAmount() {
		return ErrBalanceInsufficient
	}

	w.balance = value_objects.NewBalance(w.balance.GetAmount() - balance.GetAmount())

	return nil
}

func (w Wallet) GetId() string {
	return w.id
}
func (w Wallet) GetStatus() string {
	return w.status
}
func (w Wallet) GetEnabledAt() *time.Time {
	return w.enabled_at
}
func (w Wallet) OwnedBy() string {
	return w.owned_by
}

func (w Wallet) IsDisabled() bool {
	return w.status == "disabled"
}

func (w Wallet) GetBalance() (float64, error) {
	if w.IsDisabled() {
		return 0, ErrWalletDisable
	}

	return w.balance.GetAmount(), nil
}
