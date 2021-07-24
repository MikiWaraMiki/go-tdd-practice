package model

type IDollar interface {
	IMoney
}
type Dollar struct {
	*Money
}
