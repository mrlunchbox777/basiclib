/**
 * Author: Andrew Shoell
 * File: error_test.go
 */

package flow

import (
	"errors"
	"testing"
)

func TestMustNotWithError(t *testing.T) {
	err := errors.New("test error")

	// test that MustNot panics when passed an error
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("MustNot() did not panic")
		}
	}()
	MustNot(err)
}

func TestMustNotWithNoError(t *testing.T) {
	// test that MustNot doesn't panic when passed nil
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("MustNot() panicked")
		}
	}()
	MustNot(nil)
}
