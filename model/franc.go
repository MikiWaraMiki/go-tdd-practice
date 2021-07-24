package model

type IFranc interface {
	IMoney
	Times(multiplier int) *Franc
}
type Franc struct {
	*Money
}

func NewFranc(amount int) *Franc {
	return &Franc{
		&Money{amount: amount},
	}
}

func (this *Franc) Times(multiplier int) *Franc {
	return NewFranc(this.amount * multiplier)
}
