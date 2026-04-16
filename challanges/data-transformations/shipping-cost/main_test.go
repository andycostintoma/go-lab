package main

import "testing"

func TestCalculateTotal(t *testing.T) {
	rules := OrderRules{
		"US": {
			"mouse": []ItemRule{{0, -1, 550}},
			"laptop": []ItemRule{
				{0, 2, 1000},
				{3, 4, 950},
				{5, -1, 900},
			},
		},
	}

	order := Order{
		"US": {
			"mouse":  Quantity(20),
			"laptop": Quantity(5),
		},
	}

	got := CalculateTotal(order, rules)
	want := 15800 // 20*550 + 2*1000 + 2*950 + 1*900
	if got != want {
		t.Fatalf("expected %d, got %d", want, got)
	}
}

func TestNoRules(t *testing.T) {
	order := Order{"US": {"unknown": 5}}
	rules := OrderRules{"US": {}}

	got := CalculateTotal(order, rules)
	if got != 0 {
		t.Fatalf("expected 0, got %d", got)
	}
}

func TestProgressiveTier(t *testing.T) {
	rules := OrderRules{
		"US": {
			"book": []ItemRule{
				{0, 2, 100},
				{3, 5, 80},
				{6, -1, 60},
			},
		},
	}
	order := Order{"US": {"book": 6}}
	got := CalculateTotal(order, rules)
	want := 2*100 + 3*80 + 1*60
	if got != want {
		t.Fatalf("expected %d, got %d", want, got)
	}
}

func TestZeroQuantity(t *testing.T) {
	rules := OrderRules{"US": {"item": []ItemRule{{0, -1, 999}}}}
	order := Order{"US": {"item": 0}}
	got := CalculateTotal(order, rules)
	if got != 0 {
		t.Fatalf("expected 0, got %d", got)
	}
}
