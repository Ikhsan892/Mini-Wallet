package value_objects

type Balance struct {
	amount float64
}

func NewBalance(amount float64) Balance {
	return Balance{amount}
}

func (b Balance) GetAmount() float64 {
	return b.amount
}
