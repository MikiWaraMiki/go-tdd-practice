package model

type Expression interface {
	Reduce(string) *Money
}
