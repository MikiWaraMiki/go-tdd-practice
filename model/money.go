package model

import "fmt"

type Stringify interface {
	ToString() string
}
type IMoney interface {
	Experssion
	Equals(money IMoney) bool
	GetAmount() int
	GetCurrency() string
	Times(multiplier int) *Money
	Plus(money *Money) *Sum
}

type Money struct {
	amount   int
	currency string
}

func GenerateDollar(amount int) *Money {
	return &Money{
		amount:   amount,
		currency: "USD",
	}
}

func GenerateFranc(amount int) *Money {
	return &Money{
		amount:   amount,
		currency: "CHF",
	}
}

func (this Money) Equals(obj IMoney) bool {
	if obj == nil {
		return false
	}
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

func (this *Money) Plus(money *Money) *Sum {
	return NewSum(this, money)
}

func (this Money) ToString() string {
	return fmt.Sprintf("%v %v", this.amount, this.currency)
}
