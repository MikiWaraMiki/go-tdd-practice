package model

type IFranc interface {
	IMoney
}
type Franc struct {
	*Money
}
