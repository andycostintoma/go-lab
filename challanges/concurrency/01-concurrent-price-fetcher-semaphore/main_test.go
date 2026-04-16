package main

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func Test(t *testing.T) {
	type testCase struct {
		ids           []string
		maxConcurrent int
		ctxFactory    func() context.Context
		expectedMap   map[string]int
		expectedErr   string // substring match for readability
	}

	runCases := []testCase{
		{
			ids:           []string{"p1", "p2", "p3"},
			maxConcurrent: 2,
			ctxFactory:    func() context.Context { return context.Background() },
			expectedMap:   map[string]int{"p1": priceOf("p1"), "p2": priceOf("p2"), "p3": priceOf("p3")},
			expectedErr:   "",
		},
		{
			ids:           []string{"x", "y"},
			maxConcurrent: 1,
			ctxFactory: func() context.Context {
				ctx, cancel := context.WithCancel(context.Background())
				cancel()
				return ctx
			},
			expectedMap: nil,
			expectedErr: context.Canceled.Error(),
		},
	}

	submitCases := append(runCases, []testCase{
		{
			ids:           []string{"p1", "bad:oops", "slow:hang"},
			maxConcurrent: 3,
			ctxFactory:    func() context.Context { return context.Background() },
			expectedMap:   nil,
			expectedErr:   "fetch failed for bad:oops",
		},
		{
			ids:           []string{"slow:never", "p2"},
			maxConcurrent: 2,
			ctxFactory: func() context.Context {
				// Use an already-expired deadline to deterministically force timeout
				return expiredContext()
			},
			expectedMap: nil,
			expectedErr: context.DeadlineExceeded.Error(),
		},
		{
			ids:           []string{"alpha", "beta", "gamma", "delta", "epsilon"},
			maxConcurrent: 2,
			ctxFactory:    func() context.Context { return context.Background() },
			expectedMap: map[string]int{
				"alpha":   priceOf("alpha"),
				"beta":    priceOf("beta"),
				"gamma":   priceOf("gamma"),
				"delta":   priceOf("delta"),
				"epsilon": priceOf("epsilon"),
			},
			expectedErr: "",
		},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, tc := range testCases {
		ctx := tc.ctxFactory()
		result, err := FetchAll(ctx, tc.ids, tc.maxConcurrent)

		errStr := ""
		if err != nil {
			errStr = err.Error()
		}

		ok := false
		if tc.expectedErr == "" {
			// Expect success
			ok = err == nil && mapsEqual(tc.expectedMap, result)
		} else {
			// Expect an error containing substring
			ok = err != nil && contains(errStr, tc.expectedErr)
		}

		if !ok {
			failCount++
			t.Errorf(`---------------------------------
Input:
  IDs: %v
  maxConcurrent: %d
  Context: %T

Expected:
  Map: %v
  Error contains: %q

Actual:
  Map: %v
  Error: %v
Fail
`, tc.ids, tc.maxConcurrent, ctx, tc.expectedMap, tc.expectedErr, result, err)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Input:
  IDs: %v
  maxConcurrent: %d
  Context: %T

Expected:
  Map: %v
  Error contains: %q

Actual:
  Map: %v
  Error: %v
Pass
`, tc.ids, tc.maxConcurrent, ctx, tc.expectedMap, tc.expectedErr, result, err)
		}
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
	}
}

// Helper: deterministic expected price calculation (must mirror mockFetch success path)
func priceOf(id string) int {
	sum := 0
	for _, r := range id {
		sum += int(r)
	}
	return len(id) + (sum % 17)
}

// mapsEqual compares two string->int maps for exact match
func mapsEqual(a, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}

func contains(s, sub string) bool { return len(sub) == 0 || (len(sub) > 0 && indexOf(s, sub) >= 0) }

func indexOf(s, sub string) int {
	// simple substring search (avoid importing strings to keep test output focused)
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return i
		}
	}
	return -1
}

func expiredContext() context.Context {
	// Create a context that is already expired deterministically
	ctx, cancel := context.WithDeadline(context.Background(), time.Unix(0, 0))
	// Call cancel immediately to release resources
	cancel()
	return ctx
}

// withSubmit is set at compile time depending on which button is used to run the tests
var withSubmit = true
