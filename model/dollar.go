package model

type Dollar struct {
	amount int
}

func NewDollar(amount int) *Dollar {
	return &Dollar{
		amount: amount,
	}
}

func (this *Dollar) Times(multiplier int) *Dollar {
	return NewDollar(this.amount * multiplier)
}

func (this Dollar) Equals(obj Object) bool {
	dollar := obj.(*Dollar)
	return this.amount == dollar.amount
}
