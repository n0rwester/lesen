// Package util provides helper functions
package util

import (
	"strings"
)

func GetMaxLen(slice []string) int {
	currentMax := 0
	for _, x := range slice {
		if len(x) > currentMax {
			newlen := len(x) - len(strings.Split(x, " ")[0])
			currentMax = newlen // account for nerd font icon weirdness
		}
	}

	return currentMax
}

// Pad pads text with padding until it reaches the desired length
func Pad(text string, length int, padding string) string {
	paddingLength := length - (len(text) - 1)
	return text + strings.Repeat(padding, paddingLength)
}
