package model

import (
	"testing"
)

func TestMoneyMultiplication(t *testing.T) {
	t.Run("ドルの掛け算の結果が正しいこと", func(t *testing.T) {
		var five IDollar = GenerateDollar(5)
		if !five.Times(3).Equals(GenerateDollar(15)) {
			t.Errorf("expected: true, result: false")
		}

		if !five.Times(2).Equals(GenerateDollar(10)) {
			t.Errorf("expected: true, result: false")
		}
	})
}

func TestEquality(t *testing.T) {
	t.Run("amountが等しければ等価と判定されること", func(t *testing.T) {
		var five IDollar = NewDollar(5)

		if !five.Equals(NewDollar(5)) {
			t.Errorf("expected true, result: false")
		}

	})

	t.Run("amountが異なる場合は等価と判定されないこと", func(t *testing.T) {
		var five IDollar = NewDollar(5)

		if five.Equals(NewDollar(6)) {
			t.Error("expected: false, result: true")
		}
	})

	t.Run("ドルとフランの比較が正しいこと", func(t *testing.T) {
		var dollar IDollar = NewDollar(5)
		var franc IFranc = NewFranc(5)

		if dollar.Equals(franc) {
			t.Error("expected: false, result: true")
		}
	})
}

func TestFrancMultiplication(t *testing.T) {
	t.Run("フランの掛け算の結果が正しいこと", func(t *testing.T) {
		five := NewFranc(5)

		if !five.Times(2).Equals(NewFranc(10)) {
			t.Error("5.times(2) == 10, expected true, result: false")
		}

		if !five.Times(3).Equals(NewFranc(15)) {
			t.Error("5.times(3) == 15, expected: true, result: false")
		}
	})
}
