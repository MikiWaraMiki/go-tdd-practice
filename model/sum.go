package model

type Sum struct {
	augend IMoney
	addend IMoney
}

func NewSum(augend IMoney, addend IMoney) *Sum {
	return &Sum{
		augend: augend,
		addend: addend,
	}
}

func (this Sum) GetAugend() IMoney {
	return this.augend
}

func (this Sum) GetAddend() IMoney {
	return this.addend
}
func (this Sum) Reduce(bank *Bank, to string) *Money {
	reducedAugend := this.augend.Reduce(bank, to)
	reducedAddend := this.addend.Reduce(bank, to)

	result := reducedAugend.GetAmount() + reducedAddend.GetAmount()
	return NewMoney(result, to)
}
