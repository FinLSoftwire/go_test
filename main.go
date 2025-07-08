package main

import (
	"bufio"
	"fmt"
	"os"
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

func greeting(name string) string {
	return "Hello, " + name
}
