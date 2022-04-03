package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

func main() {
	userName := readString("What is your name?")
	userAge := readInt("How old are you?")

	// fmt.Println("Your name is"+userName+". You are age ", userAge, "years old.")
	// fmt.Println(fmt.Sprintf("Your name is %s. You are %d years old", userName, userAge))
	fmt.Printf("Your name is %s. You are %d years old.\n", userName, userAge)
}

func prompt() {
	fmt.Print("-> ")
}

func readString(s string) string {
	for {
		fmt.Println(s)
		prompt()

		reader = bufio.NewReader(os.Stdin)
		userInput, _ := reader.ReadString('\n')
		// Strip off the carriage return, first on Windows then on Mac
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		if userInput == "" {
			fmt.Println("Please enter a value")
		} else {
			return userInput
		}
	}
}

func readInt(s string) int {
	for {
		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		num, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Please enter a whole number")
		} else {
			return num
		}
	}
}
