package main

import (
	"testing"
)

func TestSubtract(t *testing.T) {
	result := Subtract(5, 3)
	expected := 2

	if result != expected {
		t.Errorf("Subtract(5, 3) = %d; want %d", result, expected)
	}
}
