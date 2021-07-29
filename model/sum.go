package model

type Sum struct {
	augend *Money
	addend *Money
}

func NewSum(augend *Money, addend *Money) *Sum {
	return &Sum{
		augend: augend,
		addend: addend,
	}
}

func (this Sum) GetAugend() *Money {
	return this.augend
}

func (this Sum) GetAddend() *Money {
	return this.addend
}
func (this Sum) Reduce(to string) *Money {
	result := this.GetAugend().GetAmount() + this.GetAddend().GetAmount()
	return NewMoney(result, to)
}
