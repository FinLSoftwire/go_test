package main

import "testing"

var tests = []struct {
	name     string
	expected string
}{
	{"Alice", "Hello, Alice"},
	{"Rob Pike", "Hello, Rob. Thanks for creating me!"},
}

func TestGreeting(t *testing.T) {
	for _, test := range tests {
		result := greeting(test.name)
		if result != test.expected {
			t.Errorf("Incorrect greeting for name %s, got: %s but expected: %s", test.name, result, test.expected)
		}
	}
}
