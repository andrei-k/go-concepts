package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What is your name?")
	prompt()

	userInput, _ := reader.ReadString('\n')

	// Strip off the carriage return (Windows)
	userInput = strings.Replace(userInput, "\r\n", "", -1)
	// (Mac)
	userInput = strings.Replace(userInput, "\n", "", -1)

	fmt.Println("Your name is", userInput)
}

func prompt() {
	fmt.Print("-> ")
}
