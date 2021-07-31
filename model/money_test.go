package model

import (
	"testing"
)

func TestMoneyMultiplication(t *testing.T) {
	t.Run("ドルの掛け算の結果が正しいこと", func(t *testing.T) {
		var five IMoney = GenerateDollar(5)
		if !five.Times(3).(IMoney).Equals(GenerateDollar(15)) {
			t.Fatalf("expected: true, result: false")
		}

		if !five.Times(2).(IMoney).Equals(GenerateDollar(10)) {
			t.Fatalf("expected: true, result: false")
		}
	})
}

func TestEquality(t *testing.T) {
	t.Run("amountが等しければ等価と判定されること", func(t *testing.T) {
		var five IMoney = GenerateDollar(5)

		if !five.Equals(GenerateDollar(5)) {
			t.Fatalf("expected true, result: false")
		}

	})

	t.Run("amountが異なる場合は等価と判定されないこと", func(t *testing.T) {
		var five IMoney = GenerateDollar(5)

		if five.Equals(GenerateDollar(6)) {
			t.Error("expected: false, result: true")
		}
	})

	t.Run("通貨単位が異なる場合は等価と判定されないこと", func(t *testing.T) {
		var dollar IMoney = GenerateDollar(5)
		var franc IMoney = GenerateFranc(5)

		if dollar.Equals(franc) {
			t.Error("expected: false, result: true")
		}
	})
	t.Run("nilが渡された場合は等価と判定されないこと", func(t *testing.T) {
		var dollar IMoney = GenerateDollar(5)

		if dollar.Equals(nil) {
			t.Error("expected: false, result: true")
		}
	})
}

func TestCurrency(t *testing.T) {
	t.Run("ドルの場合はUSDであること", func(t *testing.T) {
		dollar := GenerateDollar(5)
		if result := dollar.GetCurrency(); result != "USD" {
			t.Fatalf("expected: USD, result: %v", result)
		}
	})
	t.Run("フランの場合はCHFであること", func(t *testing.T) {
		franc := GenerateFranc(5)

		if result := franc.GetCurrency(); result != "CHF" {
			t.Fatalf("expected: CHF, result: %v", result)
		}
	})
}

func TestReduceSum(t *testing.T) {
	t.Run("足し算の結果が正しいこと", func(t *testing.T) {
		var sum Expression = NewSum(GenerateDollar(5), GenerateDollar(2))
		bank := NewBank()

		expected := GenerateDollar(7)

		if result := bank.Reduce(sum, "USD"); !result.Equals(expected) {
			t.Fatalf("result=%v, expected=%v", result.GetAmount(), expected.GetAmount())
		}
	})
}

func TestReduceMoneyDifferentCurrency(t *testing.T) {
	t.Run("通貨の変換結果が正しいこと", func(t *testing.T) {
		bank := NewBank()

		bank.AddRate("CHF", "USD", 2)

		expected := GenerateDollar(1)

		if result := bank.Reduce(GenerateFranc(2), "USD"); !expected.Equals(result) {
			t.Fatalf("expected=%v %v, result=%v %v", expected.GetAmount(), expected.GetCurrency(), result.GetAmount(), result.GetCurrency())
		}
	})
}

func TestMixedAddition(t *testing.T) {
	t.Run("異なる通貨単位同士の足し算結果が正しいこと", func(t *testing.T) {
		var fiveDollar Expression = GenerateDollar(5)
		var tenFranc Expression = GenerateFranc(10)

		bank := NewBank()

		bank.AddRate("CHF", "USD", 2)

		expected := GenerateDollar(10)
		result := bank.Reduce(fiveDollar.Plus(tenFranc), "USD")

		if !result.Equals(expected) {
			t.Fatalf("expected=%v %v, result=%v %v", expected.GetAmount(), expected.GetCurrency(), result.GetAmount(), result.GetCurrency())
		}
	})
}

func TestSumTimes(t *testing.T) {
	t.Run("掛け算の結果が正しいこと", func(t *testing.T) {
		var fiveDollar Expression = GenerateDollar(5)
		var tenFranc Expression = GenerateFranc(10)

		bank := NewBank()

		bank.AddRate("CHF", "USD", 2)

		expected := GenerateDollar(20)
		result := bank.Reduce(fiveDollar.Plus(tenFranc), "USD").Times(2).(IMoney)

		if !result.Equals(expected) {
			t.Fatalf("expected=%v %v, result=%v %v", expected.GetAmount(), expected.GetCurrency(), result.GetAmount(), result.GetCurrency())
		}
	})
}
