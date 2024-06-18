package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// StartRepl starts the REPL (Read-Eval-Print-Loop)
func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter text: ")
		scanner.Scan()

		text := scanner.Text()

		cleaned := cleanInput(text)
		// If nothing was entered, continue the loop
		if len(cleaned) == 0 {
			continue
		}


		fmt.Println("echoing", cleaned)
	}

}

//Cleans up user input and returns each word as a Token
func cleanInput(str string) []string {
	loweredStr := strings.ToLower(str)
	words := strings.Fields(loweredStr)
	return words
}