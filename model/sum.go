package model

type Sum struct {
	augend Expression
	addend Expression
}

func NewSum(augend Expression, addend Expression) *Sum {
	return &Sum{
		augend: augend,
		addend: addend,
	}
}

func (this Sum) GetAugend() Expression {
	return this.augend
}

func (this Sum) GetAddend() Expression {
	return this.addend
}
func (this Sum) Reduce(bank *Bank, to string) *Money {
	reducedAugend := this.augend.Reduce(bank, to)
	reducedAddend := this.addend.Reduce(bank, to)

	result := reducedAugend.GetAmount() + reducedAddend.GetAmount()
	return NewMoney(result, to)
}

func (this Sum) Plus(addend Expression) *Sum {
	return &Sum{}
}
