package model

import (
	"testing"
)

func TestMoneyMultiplication(t *testing.T) {
	t.Run("ドルの掛け算の結果が正しいこと", func(t *testing.T) {
		var five IMoney = GenerateDollar(5)
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
		var five IDollar = GenerateDollar(5)

		if !five.Equals(GenerateDollar(5)) {
			t.Errorf("expected true, result: false")
		}

	})

	t.Run("amountが異なる場合は等価と判定されないこと", func(t *testing.T) {
		var five IDollar = GenerateDollar(5)

		if five.Equals(GenerateDollar(6)) {
			t.Error("expected: false, result: true")
		}
	})

	t.Run("ドルとフランの比較が正しいこと", func(t *testing.T) {
		var dollar IDollar = GenerateDollar(5)
		var franc IFranc = GenerateFranc(5)

		if dollar.Equals(franc) {
			t.Error("expected: false, result: true")
		}
	})
}

func TestFrancMultiplication(t *testing.T) {
	t.Run("フランの掛け算の結果が正しいこと", func(t *testing.T) {
		five := GenerateFranc(5)

		if !five.Times(2).Equals(GenerateFranc(10)) {
			t.Error("5.times(2) == 10, expected true, result: false")
		}

		if !five.Times(3).Equals(GenerateFranc(15)) {
			t.Error("5.times(3) == 15, expected: true, result: false")
		}
	})
}

func TestCurrency(t *testing.T) {
	t.Run("ドルの場合はUSDであること", func(t *testing.T) {
		dollar := GenerateDollar(5)
		if result := dollar.GetCurrency(); result != "USD" {
			t.Errorf("expected: USD, result: %v", result)
		}
	})
	t.Run("フランの場合はCHFであること", func(t *testing.T) {
		franc := GenerateFranc(5)

		if result := franc.GetCurrency(); result != "CHF" {
			t.Errorf("expected: CHF, result: %v", result)
		}
	})
}
