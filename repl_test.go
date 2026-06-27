package main

import (
	"testing"

	"github.com/ako1993/pokedexgo/internal/utils"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur Pikachu",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "      This       has         space",
			expected: []string{"this", "has", "space"},
		},
	}
	for _, c := range cases {
		actual := utils.CleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Error actual length and expected length do not match. Test Failed")
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("%s and %s do not match. Test failed.", word, expectedWord)
			}
		}
	}
}
