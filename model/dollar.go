package model

type IDollar interface {
	IMoney
	Times(multiplier int) *Dollar
}
type Dollar struct {
	*Money
}

func NewDollar(amount int) *Dollar {
	return &Dollar{
		&Money{
			amount:   amount,
			currency: "Dollar",
		},
	}
}

func (this Dollar) Times(multiplier int) *Dollar {
	return NewDollar(this.amount * multiplier)
}

func (this Dollar) GetName() string {
	return "Dollar"
}
