package model

type IFranc interface {
	IMoney
}
type Franc struct {
	*Money
}

func NewFranc(amount int) *Franc {
	return &Franc{
		&Money{
			amount:   amount,
			currency: "Franc",
		},
	}
}
