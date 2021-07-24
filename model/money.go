package model

import "fmt"

type IMoney interface {
	Equals(money IMoney) bool
	GetAmount() int
	GetCurrency() string
	Times(multiplier int) *Money
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

func GenerateFranc(amount int) *Franc {
	return &Franc{
		&Money{
			amount:   amount,
			currency: "Franc",
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

func (this Money) Times(multiplier int) *Money {
	return &Money{
		amount:   this.amount * multiplier,
		currency: this.currency,
	}
}
