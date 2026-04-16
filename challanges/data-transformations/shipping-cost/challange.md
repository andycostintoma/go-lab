# Step 1 — Single-Country, Marginal Tiers

## Goal

Compute the total shipping cost for an order to **one country** given **per-product marginal tier pricing** (progressive per-unit cost across ranges).

## Rules

* Pricing is defined **per product** as ordered tiers: each tier has `minQuantity`, `maxQuantity` (inclusive; `-1` means no upper bound), and a per-unit `cost`.
* For each item, you **fill tiers progressively** until its full quantity is priced.
* If a product in the order has **no pricing defined** for the country, its cost is **0** (don’t error in Step 1).
* Quantities are non-negative integers. Costs are integers. Output is an integer.

## Input

* `order`:

    * `country: string`
    * `items: Array<{ name: string, quantity: number }>`
* `rules`:

    * `{ [country: string]: { [product: string]: Array<{ minQuantity: number, maxQuantity: number, cost: number }> } }`

## Output

* A single integer: the **total shipping cost**.

## Example

**Order**

```json
{
  "country": "US",
  "items": [
    { "name": "mouse", "quantity": 20 },
    { "name": "laptop", "quantity": 5 }
  ]
}
```

**Rules**

```json
{
  "US": {
    "mouse": [
      { "minQuantity": 0, "maxQuantity": -1, "cost": 550 }
    ],
    "laptop": [
      { "minQuantity": 0, "maxQuantity": 2, "cost": 1000 },
      { "minQuantity": 3, "maxQuantity": 4, "cost": 950 },
      { "minQuantity": 5, "maxQuantity": -1, "cost": 900 }
    ]
  }
}
```

**Expected**

```
15800 = 20*550 + (2*1000 + 2*950 + 1*900)
```

## Task

Implement a function that returns the total cost.

**If you want a suggested signature (optional):**

* Go:

  ```go
  func ShippingTotal(order Order, rules Rules) int
  ```
* JS/TS:

  ```ts
  function shippingTotal(order: Order, rules: Rules): number
  ```

## Assumptions to keep it simple (Step 1)

* The selected country exists in `rules` (but some products might not—treat those as cost 0).
* Tiers for a product:

    * Are non-overlapping.
    * Are ordered by `minQuantity`.
    * Start at `minQuantity = 0` and fully cover all positive quantities (last tier may have `maxQuantity = -1`).
* No currency handling, no decimals.

---

## Sample Test Cases

### Test 1 (single flat tier)

**Order**

```json
{"country":"US","items":[{"name":"sticker","quantity":3}]}
```

**Rules**

```json
{"US":{"sticker":[{"minQuantity":0,"maxQuantity":-1,"cost":10}]}}
```

**Expected**

```
30
```

### Test 2 (progressive tiers)

**Order**

```json
{"country":"US","items":[{"name":"book","quantity":6}]}
```

**Rules**

```json
{
  "US": {
    "book": [
      {"minQuantity":0,"maxQuantity":2,"cost":100},
      {"minQuantity":3,"maxQuantity":5,"cost":80},
      {"minQuantity":6,"maxQuantity":-1,"cost":60}
    ]
  }
}
```

**Expected**

```
2*100 + 3*80 + 1*60 = 500
```

### Test 3 (product missing in rules → cost 0)

**Order**

```json
{"country":"US","items":[{"name":"mouse","quantity":5}]}
```

**Rules**

```json
{"US":{"keyboard":[{"minQuantity":0,"maxQuantity":-1,"cost":200}]}}
```

**Expected**

```
0
```

### Test 4 (multiple items, mixed tiers)

**Order**

```json
{
  "country":"US",
  "items":[
    {"name":"pen","quantity":1},
    {"name":"notebook","quantity":4}
  ]
}
```

**Rules**

```json
{
  "US": {
    "pen":[{"minQuantity":0,"maxQuantity":-1,"cost":25}],
    "notebook":[
      {"minQuantity":0,"maxQuantity":1,"cost":300},
      {"minQuantity":2,"maxQuantity":3,"cost":250},
      {"minQuantity":4,"maxQuantity":-1,"cost":200}
    ]
  }
}
```

**Expected**

```
1*25 + (1*300 + 2*250 + 1*200) = 25 + 1000 = 1025
```

### Test 5 (zero quantity)

**Order**

```json
{"country":"US","items":[{"name":"widget","quantity":0}]}
```

**Rules**

```json
{"US":{"widget":[{"minQuantity":0,"maxQuantity":-1,"cost":999}]}}
```

**Expected**

```
0
```

---
