/**
 * Author: Andrew Shoell
 * File: error_test.go
 */

package flow

import (
	"errors"
	"testing"
)

// TestMustWithError tests that Must panics when passed an error
func TestMustWithError(t *testing.T) {
	err := errors.New("test error")

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Must() panicked")
		}
	}()
	Must(err)
}

// TestMustWithNoError tests that Must doesn't panic when passed nil
func TestMustWithNoError(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Must() did not panic")
		}
	}()
	Must(nil)
}

// TestMustNotWithError tests that MustNot panics when passed an error
func TestMustNotWithError(t *testing.T) {
	err := errors.New("test error")

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("MustNot() did not panic")
		}
	}()
	MustNot(err)
}

// TestMustNotWithNoError tests that MustNot doesn't panic when passed nil
func TestMustNotWithNoError(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("MustNot() panicked")
		}
	}()
	MustNot(nil)
}
