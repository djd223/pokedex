package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "hello world",
			expected: []string{
				"hello",
				"world",
			},
		},
		{
			input: "   ",
			expected: []string{},
		},
		{
			input: "HEllo World  ",
			expected: []string{
				"hello",
				"world",
			},
		},
	}
	
	for _, cs := range cases {
		actual := cleanInput(cs.input)
		if len(actual) != len(cs.expected){
			t.Errorf("The lengths are not equal: '%v' vs '%v'", actual, cs.expected)
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := cs.expected[i]
			if word != expectedWord {
				t.Errorf("cleanInput(%v) == %v, expected %v", cs.input, actual, cs.expected)
			}
		}
	}
}