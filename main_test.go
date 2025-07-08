package main

import (
	"testing"
)

var tests = []struct {
	name     string
	expected string
}{
	{"Alice", "Hello, Alice"},
	{"Alice lastname", "Hello, Alice"},
	{"Robert Griesemer", "Hello, Robert. Thanks for creating me!"},
	{"Rob Pike", "Hello, Rob. Thanks for creating me!"},
	{"Ken Thompson", "Hello, Ken. Thanks for creating me!"},
	{"racecar", "Hello, racecar. Cool, a palindromic name!"},
}

func TestGreeting(t *testing.T) {
	for _, test := range tests {
		result := greeting(test.name)
		if result != test.expected {
			t.Errorf("Incorrect greeting for name %s, got: %s but expected: %s", test.name, result, test.expected)
		}
	}
}

func TestNameInput(t *testing.T) {
	result := handleEmptyName("")
	if result != "Ok, no greeting for you" {
		t.Errorf("Incorrect response to empty name, got: %s but expected: Ok, no greeting for you", result)
	}
}
