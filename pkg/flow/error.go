/**
 * Author: Andrew Shoell
 * File: error.go
 */

package flow

import (
	"errors"
	"fmt"
)

// Must panics if err is non-nil
func Must(err error) {
	if err == nil {
		message := "original err was nil when it was expected to be non-nil"
		fmt.Printf("err: %s\n", message)
		panic(errors.New(message))
	}
}

// MustNot panics if err is nil
func MustNot(err error) {
	if err != nil {
		fmt.Printf("err: %s\n", err)
		panic(err)
	}
}
