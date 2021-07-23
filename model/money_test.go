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

	result := reflect.DeepEqual(five, NewDollar(5))
	if !result {
		t.Errorf("expected true, result: false")
	}
}
