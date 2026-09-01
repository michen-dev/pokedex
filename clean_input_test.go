package main 

import (
	"testing"
	"reflect"
)

func TestCleanInput(t *testing.T) {
	test_cases := []struct {
		input string
		expected []string
	}{
		{
		input: "I am Spiderman",
		expected: []string{"i", "am", "spiderman"},
		},
		{
		input: "  hello  world  ",
		expected: []string{"hello", "world"},
		},
		{
		input: "  nOtLowEr   Caszz ",
		expected: []string{"notlower", "caszz"},
		},
	}

	for _, c := range test_cases {
		got := cleanInput(c.input)
		if !reflect.DeepEqual(got, c.expected) {
			t.Fatalf("Got: %v\nExpected: %v", got, c.expected)
		}
	}
}