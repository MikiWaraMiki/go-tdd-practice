package model

type Object interface{}
type IMoney interface {
	Equals(money *Object) bool
}

type Money struct {
	amount int
}

func (this Money) Equals(obj Object) bool {
	money := obj.(Money)
	return this.amount == money.amount
}
