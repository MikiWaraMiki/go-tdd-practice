package model

import (
	"reflect"
	"testing"
)

func TestMoneyMultiplication(t *testing.T) {
	five := NewDollar(5)

	product := five.Times(3)

	if product.Amount() != 15 {
		t.Errorf("expected: 15, result:%v", product.Amount())
	}

	product = five.Times(2)
	if product.Amount() != 10 {
		t.Errorf("expected: 10, result: %v", product.Amount())
	}
}

func TestEquality(t *testing.T) {
	five := NewDollar(5)

	if !reflect.DeepEqual(five, NewDollar(5)) {
		t.Errorf("expected true, result: false")
	}

	if reflect.DeepEqual(five, NewDollar(6)) {
		t.Error("expected: false, result: true")
	}
}
