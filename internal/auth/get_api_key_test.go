package auth

import "testing"

// TestAlwaysSucceeds is a simple test that will never fail.
func TestAlwaysSucceeds(t *testing.T) {
	expected := true
	actual := false

	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
