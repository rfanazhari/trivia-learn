package account

type Money struct {
	amount int64
}

func NewMoney(amount int64) (Money, error) {
	if amount < 0 {
		return Money{}, ErrNegativeAmount
	}
	return Money{
		amount: amount,
	}, nil
}

func (m Money) Amount() int64 {
	return m.amount
}
