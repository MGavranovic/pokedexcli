package main

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := map[string]struct {
		input    string
		expected []string
	}{
		"simple": {
			input:    " Hello world",
			expected: []string{"hello", "world"},
		},
		"with punctuation": {
			input:    " Hello world!",
			expected: []string{"hello", "world!"},
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			actual := cleanInput(c.input)
			if !reflect.DeepEqual(c.expected, actual) {
				t.Fatalf("Input: %v |Expected: %v | Actual: %v", c.input, c.expected, actual)
			}
		})
	}
}
