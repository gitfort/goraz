package local

type Money int64

func (m Money) Int64() int64 {
	return int64(m)
}

func (m *Money) Sum(money Money) {
	*m = Money(m.Int64() + money.Int64())
}
