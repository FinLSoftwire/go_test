package main

import (
	"bytes"
	"io"
	"os"
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
	originalStdout := os.Stdout
	var buf bytes.Buffer
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("Failed to create Pipe")
	}
	os.Stdout = w
	main()
	w.Close()
	os.Stdout = originalStdout
	io.Copy(&buf, r)
	if buf.String() != "Enter your name: Ok, no greeting for you\n" {
		t.Errorf("Found: %s, expected: Ok, no greeting for you", buf.String())
	}
}
