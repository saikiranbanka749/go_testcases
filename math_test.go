package main

import (
	"testing"
)

// TestAdd is the unit test for the Add function
func TestAdd(t *testing.T) {
	result := Add(2, 3) // Calling the Add function
	expected := 5       // Expected result is 5

	if result != expected {
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	}
}
