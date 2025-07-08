package main

import "testing"

var tests = []struct {
	name     string
	expected string
}{
	{"Alice", "Hello, Alice"},
	{"Alice lastname", "Hello, Alice"},
	{"Robert Griesemer", "Hello, Robert. Thanks for creating me!"},
	{"Rob Pike", "Hello, Rob. Thanks for creating me!"},
	{"Ken Thompson", "Hello, Ken. Thanks for creating me!"},
}

func TestGreeting(t *testing.T) {
	for _, test := range tests {
		result := greeting(test.name)
		if result != test.expected {
			t.Errorf("Incorrect greeting for name %s, got: %s but expected: %s", test.name, result, test.expected)
		}
	}
}
