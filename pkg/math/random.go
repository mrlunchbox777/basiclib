/**
 * Author: Andrew Shoell
 * File: random.go
 */

package math

import (
	rand "math/rand"
)

var defaultLetterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// RandomInt returns a random integer between min and max
func RandomInt(min, max int) int {
	return rand.Intn(max-min) + min
}

// CustomRandomString returns a random string of length length using the runes provided
func CustomRandomString(length int, runes []rune) string {
	result := make([]rune, length)
	for i := range result {
		result[i] = runes[RandomInt(0, len(runes))]
	}
	return string(result)
}

// RandomString returns a random string of length length using the default runes
func RandomString(length int) string {
	return CustomRandomString(length, defaultLetterRunes)
}
