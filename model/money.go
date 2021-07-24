package model

type IMoney interface {
	Equals(money IMoney) bool
	GetAmount() int
}

type Money struct {
	amount int
}

func (this Money) Equals(obj IMoney) bool {
	return this.GetAmount() == obj.GetAmount()
}

func (this Money) GetAmount() int {
	return this.amount
}

func (this Money) GetName() string {
	return "Money"
}
