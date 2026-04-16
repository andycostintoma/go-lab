package main

import (
	"math"
)

// --- data models ---

type Item struct {
	SKU       string
	UnitPrice float64
	Qty       int
}

type ProcessingFee struct {
	Percent float64
	Fixed   float64
}

type OrderInput struct {
	Items              []Item
	VATRate            float64
	ProcessingFee      ProcessingFee
	PlatformFeePercent float64

	// Step 3 additions
	Currency      string
	ExchangeRates map[string]float64 // e.g. {"EUR": 4.97, "RON": 1.0}
}

type OrderSummary struct {
	Subtotal        float64
	VAT             float64
	TotalCollected  float64
	ProcessingFee   float64
	PlatformFee     float64
	SellerPayout    float64
	PlatformRevenue float64

	// Step 3 additions
	Currency           string
	SellerPayoutRON    float64
	PlatformRevenueRON float64
}

// --- main logic ---

func CalculateOrderSummary(in OrderInput) OrderSummary {
	// --- compute subtotal with validation ---
	subtotal := 0.0
	for _, item := range in.Items {
		if item.Qty > 0 && item.UnitPrice >= 0 {
			subtotal += item.UnitPrice * float64(item.Qty)
		}
	}

	if subtotal == 0 {
		return OrderSummary{}
	}

	// --- normalize inputs ---
	vatRate := defaultIf(in.VATRate, 0.19, func(v float64) bool { return v <= 0 })
	pfPercent := clampMin(in.ProcessingFee.Percent, 0)
	pfFixed := clampMin(in.ProcessingFee.Fixed, 0)
	platformFeePercent := clamp(in.PlatformFeePercent, 0, 1)

	// --- main arithmetic (in order currency) ---
	vat := subtotal * vatRate
	total := subtotal + vat
	processingFee := total*pfPercent + pfFixed
	platformFee := subtotal * platformFeePercent
	sellerPayout := total - processingFee - platformFee - vat

	// --- currency rounding ---
	subtotal = roundCurrency(subtotal, in.Currency)
	vat = roundCurrency(vat, in.Currency)
	total = roundCurrency(total, in.Currency)
	processingFee = roundCurrency(processingFee, in.Currency)
	platformFee = roundCurrency(platformFee, in.Currency)
	sellerPayout = roundCurrency(sellerPayout, in.Currency)

	// --- conversion to RON ---
	rate := 1.0
	if in.ExchangeRates != nil {
		if r, ok := in.ExchangeRates[in.Currency]; ok && r > 0 {
			rate = r
		}
	}
	sellerPayoutRON := round2(roundCurrency(sellerPayout, in.Currency) * rate)
	platformRevenueRON := round2(roundCurrency(platformFee, in.Currency) * rate)

	// --- return summary ---
	return OrderSummary{
		Currency:           in.Currency,
		Subtotal:           subtotal,
		VAT:                vat,
		TotalCollected:     total,
		ProcessingFee:      processingFee,
		PlatformFee:        platformFee,
		SellerPayout:       sellerPayout,
		PlatformRevenue:    platformFee,
		SellerPayoutRON:    round2(sellerPayoutRON),
		PlatformRevenueRON: round2(platformRevenueRON),
	}
}

// --- helpers ---

func round2(v float64) float64 {
	return math.Round(v*100) / 100
}

func roundCurrency(v float64, currency string) float64 {
	if currency == "JPY" {
		return math.Round(v) // 0 decimals
	}
	return round2(v) // 2 decimals
}

func clamp(v, min, max float64) float64 {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}

func clampMin(v, min float64) float64 {
	if v < min {
		return min
	}
	return v
}

func defaultIf(v, def float64, invalid func(float64) bool) float64 {
	if invalid(v) {
		return def
	}
	return v
}
