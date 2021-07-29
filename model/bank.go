package model

type Bank struct {
	rates map[Pair]int
}

func NewBank() *Bank {
	return &Bank{
		rates: make(map[Pair]int),
	}
}

func (this Bank) Reduce(source Expression, to string) *Money {
	return source.Reduce(&this, to)
}

func (this *Bank) AddRate(from string, to string, rate int) {
	pair := NewPair(from, to)
	this.rates[pair] = rate
}

func (this Bank) GetRate(from string, to string) int {
	if from == to {
		return 1
	}
	pair := NewPair(from, to)
	return this.rates[pair]
}
