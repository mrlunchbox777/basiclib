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

// TestReadArbitraryFilePanic tests that ReadFile panics on error
func TestReadArbitraryFilePanic(t *testing.T) {
	Cleanup()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("ReadArbitraryFile(%s) should have panicked", path)
		}
	}()

	ReadArbitraryFile(path)
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

// TestReadTextFilePanic tests that ReadTextFile panics on error
func TestReadTextFilePanic(t *testing.T) {
	Cleanup()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("ReadTextFile(%s) should have panicked", path)
		}
	}()

	ReadTextFile(path)
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

// TestReadFilePanic tests that ReadFile panics on error
func TestReadFilePanic(t *testing.T) {
	Cleanup()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("ReadFile(%s) should have panicked", path)
		}
	}()

	ReadFile(path)
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

// TestWriteFilePanic tests that WriteFile panics on error
func TestWriteFilePanic(t *testing.T) {
	Cleanup()
	badPath := "/this/is/a/bad/path"
	content := math.RandomString(10)
	perms := 0644

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("WriteFile(%s, %s, %d) should have panicked", badPath, content, perms)
		}
	}()

	WriteFile(badPath, []byte(content), os.FileMode(perms))
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

// TestRemovePanic tests that Remove panics on error
func TestRemovePanic(t *testing.T) {
	Cleanup()
	badPath := "/this/is/a/bad/path"

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Remove(%s) should have panicked", badPath)
		}
	}()

	Remove(badPath)
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
