package control_structures

import "testing"

func TestCollatzChainLength(t *testing.T) {
	n := collatzChainLength(27)

	if n != 111 {
		t.Errorf("Expected %d, got %d", 111, n)
	}
}
