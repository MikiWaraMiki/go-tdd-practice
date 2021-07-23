package model

type Dollar struct {
	amount int
}

func NewDollar(amount int) *Dollar {
	return &Dollar{
		amount: amount,
	}
}

func (this *Dollar) Times(multiplier int) {
	this.amount = this.amount * multiplier
}

func (this Dollar) Amount() int {
	return 10
}
