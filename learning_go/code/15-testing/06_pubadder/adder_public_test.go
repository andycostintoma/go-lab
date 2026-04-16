package pubadder_test

import (
	"testing"
)

func TestAddNumbers(t *testing.T) {
	result := AddNumbers(2, 3)
	if result != 5 {
		t.Error("incorrect result: expected 5, got", result)
	}
}
