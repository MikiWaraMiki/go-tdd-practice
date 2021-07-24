package model

type IDollar interface {
	IMoney
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

func (this Dollar) GetName() string {
	return "Dollar"
}
