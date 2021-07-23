package model

type IDollar struct {
	IMoney
}
type Dollar struct {
	Money
}

func NewDollar(amount int) *Dollar {
	return &Dollar{
		Money{amount: amount},
	}
}

func (this *Dollar) Times(multiplier int) *Dollar {
	return NewDollar(this.amount * multiplier)
}
