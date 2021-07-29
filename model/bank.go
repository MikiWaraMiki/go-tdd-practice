package model

type Bank struct {
}

func NewBank() *Bank {
	return &Bank{}
}

func (this Bank) Reduce(source Expression, to string) *Money {
	return source.Reduce(to)
}

func (this Bank) AddRate(from string, to string, rate int) {

}
