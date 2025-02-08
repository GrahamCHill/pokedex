package tests

import (
	"github.com/grahamchill/pokedex/functions"
	"testing"
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
			input:    "  Apple  Sauce  is  Wonderful  ",
			expected: []string{"Apple", "Sauce", "is", "Wonderful"},
		},
		{
			input:    "  Mango  relish   ",
			expected: []string{"Mango", "relish"},
		},
	}

	for _, c := range cases {
		actual := functions.CleanInput(c.input)
		// Check the length of the actual slice
		if len(actual) != len(c.expected) {
			t.Errorf("Length does not match. Expected: %d, Actual: %d", len(c.expected), len(actual))
		}
		// Check each word in the slice
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Words do not match. Expected: %v, Actual: %v", expectedWord, word)
			}
		}
	}
}
