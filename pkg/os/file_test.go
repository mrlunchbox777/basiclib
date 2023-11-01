/**
 * Author: Andrew Shoell
 * File: file_test.go
 */

package os

import (
	"os"
	"testing"

	math "github.com/mrlunchbox777/basiclib/pkg/math"
)

var path = "test.txt"

func Cleanup() {
	if Exists(path) {
		Remove(path)
	}
}

// TestReadArbitraryFile tests that ReadFile returns the contents of a file
func TestReadArbitraryFile(t *testing.T) {
	Cleanup()
	content := math.RandomString(10)
	perms := 0644

	WriteFile(path, []byte(content), os.FileMode(perms))

	if got := string(ReadArbitraryFile(path)); got != content {
		t.Errorf("ReadArbitraryFile(%s) = %s, didn't return %s", path, got, content)
	}
	Cleanup()
}

// TestReadTextFile tests that ReadTextFile returns the contents of a text file
func TestReadTextFile(t *testing.T) {
	Cleanup()
	content := math.RandomString(10)
	perms := 0644

	WriteFile(path, []byte(content), os.FileMode(perms))

	if got := ReadTextFile(path); got != content {
		t.Errorf("ReadTextFile(%s) = %s, didn't return %s", path, got, content)
	}
	Cleanup()
}

// TestReadFile tests that ReadFile returns the contents of a text file
func TestReadFile(t *testing.T) {
	Cleanup()
	content := math.RandomString(10)
	perms := 0644

	WriteFile(path, []byte(content), os.FileMode(perms))

	if got := ReadFile(path); got != content {
		t.Errorf("ReadFile(%s) = %s, didn't return %s", path, got, content)
	}
	Cleanup()
}

// TestWriteFile tests that WriteFile writes data to a file
func TestWriteFile(t *testing.T) {
	Cleanup()
	content := math.RandomString(10)
	perms := 0644

	WriteFile(path, []byte(content), os.FileMode(perms))

	if got := string(ReadArbitraryFile(path)); got != content {
		t.Errorf("WriteFile(%s, %s, %d) = %s, didn't return %s", path, content, perms, got, content)
	}
	Cleanup()
}

// Test Remove tests that Remove removes a file
func TestRemove(t *testing.T) {
	Cleanup()
	content := math.RandomString(10)
	perms := 0644

	WriteFile(path, []byte(content), os.FileMode(perms))
	Remove(path)
	_, err := os.Stat(path)
	if !os.IsNotExist(err) {
		t.Errorf("Remove(%s) didn't remove the file", path)
	}
}

// TestExists tests that Exists returns true if a file exists
func TestExists(t *testing.T) {
	Cleanup()
	content := math.RandomString(10)
	perms := 0644
	expected := false

	if got := Exists(path); got != expected {
		t.Errorf("Exists(%s) = %t, didn't return %t", path, got, expected)
	}

	WriteFile(path, []byte(content), os.FileMode(perms))
	expected = true

	if got := Exists(path); got != expected {
		t.Errorf("Exists(%s) = %t, didn't return %t", path, got, expected)
	}
}
