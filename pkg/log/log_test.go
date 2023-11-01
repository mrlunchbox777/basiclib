/**
 * Author: Andrew Shoell
 * File: log_test.go
 */

package log

import (
	"testing"

	"github.com/mrlunchbox777/basiclib/pkg/os"
)

// TestLogf tests that Logf logs a message to stdout
func TestLogf(t *testing.T) {
	expected := "Hello, World!"
	got, err := os.CaptureOutput(func() error {
		Logf(expected)
		return nil
	})

	if err != nil {
		t.Errorf("Logf() returned an error: %s", err)
	}
	if got != expected {
		t.Errorf("Logf() = \"%s\", didn't return \"%s\"", got, expected)
	}
}
