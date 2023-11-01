/**
 * Author: Andrew Shoell
 * File: output_test.go
 */

package os

import (
	"fmt"
	"testing"
)

// TestCaptureOutput tests that captureOutput captures the output of a function
func TestCaptureOutput(t *testing.T) {
	expected := "Hello, World!"
	got, err := CaptureOutput(func() error {
		fmt.Print(expected)
		return nil
	})

	if err != nil {
		t.Errorf("captureOutput() returned an error: %s", err)
	}
	if got != expected {
		t.Errorf("captureOutput() = \"%s\", didn't return \"%s\"", got, expected)
	}
}
