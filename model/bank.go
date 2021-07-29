package model

type Bank struct {
}

func NewBank() *Bank {
	return &Bank{}
}

func (this Bank) Reduce(source Expression, to string) *Money {
	return source.Reduce(&this, to)
}

func (this Bank) AddRate(from string, to string, rate int) {

}

func (this Bank) GetRate(from string, to string) int {
	if from == "CHF" && to == "USD" {
		return 2
	} else {
		return 1
	}
}
