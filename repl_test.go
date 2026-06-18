package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "     hello world    ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  some text   here ",
			expected: []string{"some", "text", "here"},
		},
		{
			input:    "THIS ONE IN CAPS",
			expected: []string{"this", "one", "in", "caps"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Error("Length of output different than expected")
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("%v does not match %v", word, expectedWord)
			}
		}
	}

}
