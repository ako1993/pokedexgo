package utils

import (
	"strings"
)

func CleanInput(text string) []string {
	var results []string
	words := strings.Fields(text)
	for _, word := range words {
		word = strings.ToLower(word)
		results = append(results, word)
	}
	return results
}
