/**
 * Author: Andrew Shoell
 * File: log.go
 */

package os

import (
	"io"
	"os"
)

// CaptureOutput captures the output of a function
func CaptureOutput(f func() error) (string, error) {
	orig := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	err := f()
	os.Stdout = orig
	w.Close()
	out, _ := io.ReadAll(r)
	return string(out), err
}
