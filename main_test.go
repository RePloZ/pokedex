package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	type testCase struct {
		input    string
		expected []string
	}
	cases := []testCase{
		{
			input:    "   hello    world   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf(`
---------------------------------
Test Failed:
  input: %v
  expected: %v
  actual: %v
`,
					c.input,
					c.expected,
					actual)

			}
		}
	}
}
