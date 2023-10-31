/**
 * Author: Andrew Shoell
 * File: basiclib_test.go
 */

package pkg

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	expected := "HelloWorld!"

	if got := HelloWorld(); got != expected {
		t.Errorf("HelloWorld() = %s, didn't return %s", got, expected)
	}
}
