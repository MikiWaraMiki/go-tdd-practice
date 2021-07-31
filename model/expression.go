package model

type Expression interface {
	Reduce(*Bank, string) *Money
	Plus(addend Expression) *Sum
	Times(multiplier int) Expression
}
