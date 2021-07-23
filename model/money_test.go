package model

import (
	"testing"
)

func TestMoneyMultiplication(t *testing.T) {
	t.Run("掛け算の結果が正しいこと", func(t *testing.T) {
		five := NewDollar(5)

		product := five.Times(3)

		if product.Amount() != 15 {
			t.Errorf("expected: 15, result:%v", product.Amount())
		}

		product = five.Times(2)
		if product.Amount() != 10 {
			t.Errorf("expected: 10, result: %v", product.Amount())
		}
	})
}

func TestEquality(t *testing.T) {
	t.Run("amountが等しければ等価と判定されること", func(t *testing.T) {
		five := NewDollar(5)

		if !five.Equals(NewDollar(5)) {
			t.Errorf("expected true, result: false")
		}

	})

	t.Run("amountが異なる場合は等価と判定されないこと", func(t *testing.T) {
		five := NewDollar(5)

		if five.Equals(NewDollar(6)) {
			t.Error("expected: false, result: true")
		}
	})

}
