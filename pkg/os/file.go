/**
 * Author: Andrew Shoell
 * File: file.go
 */

package os

import (
	"os"

	flow "github.com/mrlunchbox777/basiclib/pkg/flow"
)

// ReadArbitraryFile reads a file and returns its contents, panics on error
func ReadArbitraryFile(path string) []byte {
	d, err := os.ReadFile(path)
	flow.MustNot(err)
	return d
}

// ReadTextFile reads a text file and returns its contents, panics on error
func ReadTextFile(path string) string {
	d, err := os.ReadFile(path)
	flow.MustNot(err)
	return string(d)
}

// ReadFile reads a text file and returns its contents, panics on error
func ReadFile(path string) string {
	return ReadTextFile(path)
}

// WriteFile writes data to a file, panics on error
func WriteFile(path string, data []byte, perm os.FileMode) {
	err := os.WriteFile(path, data, perm)
	flow.MustNot(err)
}

// Remove removes a file or (empty) directory, panics on error
func Remove(path string) {
	err := os.Remove(path)
	flow.MustNot(err)
}

// Exists checks if a file or directory exists
func Exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
