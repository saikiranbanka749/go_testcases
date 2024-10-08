package main

import (
	"testing"
)

func TestConcat(t *testing.T) {
	result := Concat("Hello, ", "World!")
	expected := "Hello, World!"

	if result != expected {
		t.Errorf("Concat('Hello, ', 'World!') = %s; want %s", result, expected)
	}
}
