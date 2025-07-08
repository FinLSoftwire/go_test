package main

import "testing"

func TestGreeting(t *testing.T) {
	var result, expected string
	result = greeting()
	expected = "Hello, world."
	if result != expected {
		t.Errorf("incorrect greeting, got: %s expected: %s", result, expected)
	}
}
