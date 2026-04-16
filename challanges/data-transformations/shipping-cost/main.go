package main

type Name string
type Quantity int
type Country string

type Items map[Name]Quantity
type Order map[Country]Items

type ItemRule struct {
	minQuantity int
	maxQuantity int
	cost        int
}
type ItemRules map[Name][]ItemRule
type OrderRules map[Country]ItemRules

func CalculateTotal(order Order, rules OrderRules) int {
	total := 0
	for country, items := range order {
		for name, quantity := range items {
			itemRules := rules[country][name]
			total += calculateCostPerItem(itemRules, quantity)
		}
	}
	return total
}

func calculateCostPerItem(rules []ItemRule, quantity Quantity) int {
	q, total := int(quantity), 0
	for _, r := range rules {
		if q <= 0 {
			break
		}
		low := r.minQuantity
		if low == 0 {
			low = 1
		}
		high := r.maxQuantity
		if high == -1 || high > q {
			high = q
		}
		total += (high - low + 1) * r.cost

	}
	return total
}
