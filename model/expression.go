package model

type Expression interface {
	Reduce(*Bank, string) *Money
}
