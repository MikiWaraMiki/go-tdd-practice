package model

import "fmt"

type IMoney interface {
	Equals(money IMoney) bool
	GetAmount() int
	GetName() string
}

type Money struct {
	amount int
	name   string
}

func (this Money) Equals(obj IMoney) bool {
	fmt.Printf("this: %v, obj: %v\n", this.GetName(), obj.GetName())
	if this.GetName() != obj.GetName() {
		return false
	}
	return this.GetAmount() == obj.GetAmount()
}

func (this Money) GetAmount() int {
	return this.amount
}

func (this Money) GetName() string {
	return this.name
}
