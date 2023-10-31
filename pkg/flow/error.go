/**
 * Author: Andrew Shoell
 * File: error.go
 */

package flow

import (
	"fmt"
)

func MustNot(err error) {
	if err != nil {
		fmt.Printf("err: %s\n", err)
		panic(err)
	}
}
