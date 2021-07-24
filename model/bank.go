package model

type Bank struct {
}

func NewBank() *Bank {
	return &Bank{}
}

func (this Bank) Reduce(source Experssion, to string) *Money {
	return GenerateDollar(10)
}
