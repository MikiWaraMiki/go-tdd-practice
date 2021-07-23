package model

type Franc struct {
	amount int
}

func NewFranc(amount int) *Franc {
	return &Franc{
		amount: amount,
	}
}

func (this *Franc) Times(multiplier int) *Franc {
	return NewFranc(this.amount * multiplier)
}

func (this Franc) Equals(obj Object) bool {
	franc := obj.(*Franc)
	return this.amount == franc.amount
}
