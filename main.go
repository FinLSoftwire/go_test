package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	name := getName()
	fmt.Println(greeting(name))
}

func getName() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your name: ")
	scanner.Scan()
	return scanner.Text()
}

func truncateLongNames(nameString string) string {
	if len(nameString) > 20 {
		return nameString[:20] + "... Wow, that name's too long for me!"
	}
	return nameString
}

func greeting(name string) string {
	var greetingString string
	nameString := name
	spaceIndex := strings.Index(name, " ")
	if spaceIndex != -1 {
		nameString = name[:spaceIndex]
	}
	nameString = truncateLongNames(nameString)
	greetingString = "Hello, " + nameString
	creators := [3]string{"Robert Griesemer", "Rob Pike", "Ken Thompson"}
	for i := 0; i < len(creators); i++ {
		if name == creators[i] {
			greetingString = greetingString + ". Thanks for creating me!"
			break
		}
	}
	return greetingString
}
