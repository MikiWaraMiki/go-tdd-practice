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
	this.amount *= multiplier
	return nil
}

func (this Dollar) Amount() int {
	return this.amount
}
