package main

import (
	"testing"
)

// writing tests in go is just adding files ending with XX_test.go
// To run test, simply do
// > go test

// To test a function, write a function named
// Test{FunctionName} giving it the test handler as parameter as
// t *testing.T
func TestNewDeck(t *testing.T) {
	d := newDeck()
	// Test out a condition you want to be true
	// If the test failed,
	// let the test handler t log out an error as a formatted string with the function Errorf
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}
}
