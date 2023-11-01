/**
 * Author: Andrew Shoell
 * File: log.go
 */

package log

import (
	"fmt"
)

// Logf logs an optionally formatted message to stdout
func Logf(format string, args ...interface{}) {
	if len(args) == 0 {
		fmt.Print(format)
		return
	}
	fmt.Printf(format, args...)
}
