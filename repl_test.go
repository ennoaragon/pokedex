package main

import "testing"

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
			input:    "testing world ",
			expected: []string{"testing", "world"},
		},
		{
			input:    "final-as attempt here8209 18",
			expected: []string{"final-as", "attempt", "here8209", "18"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Actual len: %d and expected len: %d, not equal", len(actual), len(c.expected))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Words are not the same expected: %s got %s", expectedWord, word)
			}
		}
	}
}
