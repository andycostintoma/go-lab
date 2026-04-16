package main

import (
	"math"
	"testing"
)

func TestCalculateOrderSummary_Basic(t *testing.T) {
	in := OrderInput{
		Items: []Item{
			{SKU: "TSHIRT", UnitPrice: 100.00, Qty: 2},
			{SKU: "MUG", UnitPrice: 40.00, Qty: 1},
		},
		VATRate: 0.19,
		ProcessingFee: ProcessingFee{
			Percent: 0.029,
			Fixed:   1.20,
		},
		PlatformFeePercent: 0.10,
		Currency:           "RON",
		ExchangeRates:      map[string]float64{"RON": 1.0},
	}

	got := CalculateOrderSummary(in)
	want := OrderSummary{
		Currency:           "RON",
		Subtotal:           240.00,
		VAT:                45.60,
		TotalCollected:     285.60,
		ProcessingFee:      9.48,
		PlatformFee:        24.00,
		SellerPayout:       206.52,
		PlatformRevenue:    24.00,
		SellerPayoutRON:    206.52,
		PlatformRevenueRON: 24.00,
	}

	assertEqual(t, got, want)
}

func TestCalculateOrderSummary_EmptyOrder(t *testing.T) {
	in := OrderInput{
		Items:   []Item{},
		VATRate: 0.19,
		ProcessingFee: ProcessingFee{
			Percent: 0.029,
			Fixed:   1.20,
		},
		PlatformFeePercent: 0.10,
		Currency:           "RON",
		ExchangeRates:      map[string]float64{"RON": 1.0},
	}

	got := CalculateOrderSummary(in)
	want := OrderSummary{} // all zeros
	assertEqual(t, got, want)
}

// --- Step 2 validation tests ---

func TestCalculateOrderSummary_InvalidItemsAndDefaults(t *testing.T) {
	in := OrderInput{
		Items: []Item{
			{SKU: "TSHIRT", UnitPrice: 100.00, Qty: 2},
			{SKU: "BROKEN", UnitPrice: -5.00, Qty: 3}, // ignore
			{SKU: "MUG", UnitPrice: 40.00, Qty: 0},    // ignore
		},
		VATRate: -1, // invalid → default 0.19
		ProcessingFee: ProcessingFee{
			Percent: 0.029,
			Fixed:   -5.0, // invalid → treat as 0
		},
		PlatformFeePercent: 1.2, // clamp to 1.0
		Currency:           "RON",
		ExchangeRates:      map[string]float64{"RON": 1.0},
	}

	got := CalculateOrderSummary(in)
	want := OrderSummary{
		Currency:           "RON",
		Subtotal:           200.00,
		VAT:                38.00,
		TotalCollected:     238.00,
		ProcessingFee:      6.90,
		PlatformFee:        200.00,
		SellerPayout:       -6.90,
		PlatformRevenue:    200.00,
		SellerPayoutRON:    -6.90,
		PlatformRevenueRON: 200.00,
	}

	assertEqual(t, got, want)
}

func TestCalculateOrderSummary_AllInvalidItems(t *testing.T) {
	in := OrderInput{
		Items: []Item{
			{SKU: "BAD1", UnitPrice: -10, Qty: 2},
			{SKU: "BAD2", UnitPrice: 10, Qty: 0},
		},
		VATRate: 0.19,
		ProcessingFee: ProcessingFee{
			Percent: 0.02,
			Fixed:   1.00,
		},
		PlatformFeePercent: 0.1,
		Currency:           "RON",
		ExchangeRates:      map[string]float64{"RON": 1.0},
	}

	got := CalculateOrderSummary(in)
	want := OrderSummary{}
	assertEqual(t, got, want)
}

func TestCalculateOrderSummary_NegativeFeesAndVAT(t *testing.T) {
	in := OrderInput{
		Items: []Item{
			{SKU: "GOOD", UnitPrice: 50, Qty: 1},
		},
		VATRate: -0.05, // invalid → default 0.19
		ProcessingFee: ProcessingFee{
			Percent: -0.5, // invalid → 0
			Fixed:   -2.0, // invalid → 0
		},
		PlatformFeePercent: -0.2, // clamp to 0
		Currency:           "RON",
		ExchangeRates:      map[string]float64{"RON": 1.0},
	}

	got := CalculateOrderSummary(in)
	want := OrderSummary{
		Currency:           "RON",
		Subtotal:           50.00,
		VAT:                9.50,
		TotalCollected:     59.50,
		ProcessingFee:      0.00,
		PlatformFee:        0.00,
		SellerPayout:       50.00,
		PlatformRevenue:    0.00,
		SellerPayoutRON:    50.00,
		PlatformRevenueRON: 0.00,
	}

	assertEqual(t, got, want)
}

// --- Step 3 multi-currency tests ---

func TestCalculateOrderSummary_MultiCurrencyEUR(t *testing.T) {
	in := OrderInput{
		Currency: "EUR",
		ExchangeRates: map[string]float64{
			"EUR": 4.97,
			"RON": 1.0,
		},
		Items: []Item{
			{SKU: "TSHIRT", UnitPrice: 50.00, Qty: 2},
			{SKU: "MUG", UnitPrice: 10.00, Qty: 1},
		},
		VATRate: 0.19,
		ProcessingFee: ProcessingFee{
			Percent: 0.029,
			Fixed:   0.25,
		},
		PlatformFeePercent: 0.10,
	}

	got := CalculateOrderSummary(in)
	want := OrderSummary{
		Currency:           "EUR",
		Subtotal:           110.00,
		VAT:                20.90,
		TotalCollected:     130.90,
		ProcessingFee:      4.05,
		PlatformFee:        11.00,
		SellerPayout:       94.95,
		PlatformRevenue:    11.00,
		SellerPayoutRON:    472.93,
		PlatformRevenueRON: 54.67,
	}

	assertEqual(t, got, want)
}

func TestCalculateOrderSummary_MissingExchangeRate(t *testing.T) {
	in := OrderInput{
		Currency: "USD",
		ExchangeRates: map[string]float64{
			"RON": 1.0,
		},
		Items: []Item{
			{SKU: "BOOK", UnitPrice: 20.00, Qty: 1},
		},
		VATRate: 0.10,
		ProcessingFee: ProcessingFee{
			Percent: 0.05,
			Fixed:   1.0,
		},
		PlatformFeePercent: 0.10,
	}

	got := CalculateOrderSummary(in)
	want := OrderSummary{
		Currency:           "USD",
		Subtotal:           20.00,
		VAT:                2.00,
		TotalCollected:     22.00,
		ProcessingFee:      2.10,
		PlatformFee:        2.00,
		SellerPayout:       15.90,
		PlatformRevenue:    2.00,
		SellerPayoutRON:    15.90, // since missing USD rate → assume 1.0
		PlatformRevenueRON: 2.00,
	}

	assertEqual(t, got, want)
}

func assertEqual(t *testing.T, got, want OrderSummary) {
	t.Helper()

	const eps = 0.01
	diff := func(a, b float64) bool { return math.Abs(a-b) < eps }

	if got.Currency != want.Currency ||
		!diff(got.Subtotal, want.Subtotal) ||
		!diff(got.VAT, want.VAT) ||
		!diff(got.TotalCollected, want.TotalCollected) ||
		!diff(got.ProcessingFee, want.ProcessingFee) ||
		!diff(got.PlatformFee, want.PlatformFee) ||
		!diff(got.SellerPayout, want.SellerPayout) ||
		!diff(got.PlatformRevenue, want.PlatformRevenue) ||
		!diff(got.SellerPayoutRON, want.SellerPayoutRON) ||
		!diff(got.PlatformRevenueRON, want.PlatformRevenueRON) {
		t.Errorf("\nGot:  %+v\nWant: %+v", got, want)
	}
}
