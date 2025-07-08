package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	name := getName()
	if !handleEmptyName(name) {
		fmt.Println(greeting(name))
	}
}

func getName() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your name: ")
	scanner.Scan()
	return scanner.Text()
}

func handleEmptyName(name string) bool {
	if len(name) == 0 {
		fmt.Println("Ok, no greeting for you")
		return true
	}
	return false
}

func greeting(name string) string {
	var greetingString string
	nameString := getFirstName(name)
	nameString = truncateLongNames(nameString)
	greetingString = "Hello, " + nameString
	greetingString = thankGoCreators(greetingString, name)
	greetingString = processPalindromicName(greetingString, name)
	return greetingString
}

func getFirstName(nameString string) string {
	spaceIndex := strings.Index(nameString, " ")
	if spaceIndex == -1 {
		return nameString
	}
	return nameString[:spaceIndex]
}

func truncateLongNames(nameString string) string {
	if len(nameString) > 20 {
		return nameString[:20] + "... Wow, that name's too long for me!"
	}
	return nameString
}

func thankGoCreators(currentGreetingString string, name string) string {
	creators := [3]string{"Robert Griesemer", "Rob Pike", "Ken Thompson"}
	for i := 0; i < len(creators); i++ {
		if name == creators[i] {
			return addNecessaryPunctuation(currentGreetingString) + "Thanks for creating me!"
		}
	}
	return currentGreetingString
}

func reverseString(inp string) (reversed string) {
	for _, inputCharacter := range inp {
		reversed = string(inputCharacter) + reversed
	}
	return
}

// When adding to a greeting string that does not currently end in a space, assume it needs punctuation
func addNecessaryPunctuation(sentence string) string {
	if sentence[len(sentence)-1] != ' ' {
		return sentence + ". "
	}
	return sentence
}

func processPalindromicName(currentGreetingString string, name string) string {
	if name == reverseString(name) {
		return addNecessaryPunctuation(currentGreetingString) + "Cool, a palindromic name!"
	}
	return currentGreetingString
}
