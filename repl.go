package main

import (
	"strings"
)

func cleanInput(text string) []string {
	var results []string
	words := strings.Fields(text)
	for _, word := range words {
		results = append(results, word)
	}
	return results
}
