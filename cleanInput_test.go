package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " Hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    " Hello world!",
			expected: []string{"hello", "world!"},
		},
		{
			input:    " Hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    " Hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    " Hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    " Hello world",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		t.Logf("input: %q, expected: %v, actual: %v", c.input, c.expected, actual)
		if len(actual) != len(c.expected) {
			t.Errorf("Expected %v, but got %v\n", c.expected, actual)
			t.Fail()
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Expected %v, but got %v\n", c.expected, actual)
				t.Fail()
			}
		}
	}
}
