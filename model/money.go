package model

import "fmt"

type IMoney interface {
	Equals(money IMoney) bool
	GetAmount() int
	GetCurrency() string
}

type Money struct {
	amount   int
	currency string
}

func GenerateDollar(amount int) *Dollar {
	return &Dollar{
		&Money{
			amount:   amount,
			currency: "Dollar",
		},
	}
}

func (this Money) Equals(obj IMoney) bool {
	fmt.Printf("this: %v, obj: %v\n", this.GetCurrency(), obj.GetCurrency())
	if this.GetCurrency() != obj.GetCurrency() {
		return false
	}
	return this.GetAmount() == obj.GetAmount()
}

func (this Money) GetAmount() int {
	return this.amount
}

func (this Money) GetCurrency() string {
	return this.currency
}
