package model

import "testing"

func TestMoneyMultiplication(t *testing.T) {
	five := NewDollar(5)

	five.Times(2)

	if five.Amount() != 10 {
		t.Errorf("expected: 10, result: %v", five.Amount())
	}
}
