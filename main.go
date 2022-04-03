package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader *bufio.Reader

func main() {
	userName := readString("What is your name?")
	fmt.Println("Your name is", userName)
}

func prompt() {
	fmt.Print("-> ")
}

func readString(s string) string {
	fmt.Println(s)
	prompt()

	reader = bufio.NewReader(os.Stdin)
	userInput, _ := reader.ReadString('\n')
	// Strip off the carriage return (Windows)
	userInput = strings.Replace(userInput, "\r\n", "", -1)
	// (Mac)
	userInput = strings.Replace(userInput, "\n", "", -1)

	return userInput
}
