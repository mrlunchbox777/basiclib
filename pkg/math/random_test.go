/**
 * Author: Andrew Shoell
 * File: random_test.go
 */

package math

import (
	"testing"
)

// TestRandomInt tests that RandomInt returns a random integer between min and max
func TestRandomInt(t *testing.T) {
	min := 0
	max := 10
	got := RandomInt(min, max)
	if got < min || got > max {
		t.Errorf("RandomInt(%d, %d) = %d, didn't return a number between %d and %d", min, max, got, min, max)
	}
}

// TestCustomRandomString tests that CustomRandomString returns a random string of length length using the runes provided
func TestCustomRandomString(t *testing.T) {
	runes := []rune("abc")
	length := 10
	got := CustomRandomString(length, runes)
	if len(got) != length {
		t.Errorf("CustomRandomString(%d, %q) = %q, didn't return a string of length %d", length, runes, got, length)
	}
}

// TestRandomString tests that RandomString returns a random string of length length using the default runes
func TestRandomString(t *testing.T) {
	length := 10
	got := RandomString(length)
	if len(got) != length {
		t.Errorf("RandomString(%d) = %q, didn't return a string of length %d", length, got, length)
	}
}
