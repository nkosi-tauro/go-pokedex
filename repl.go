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
		//The first inputed word is the command
		commandName := cleaned[0]

		availableCommands := getCommands()
		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("Invalid Command")
			continue
		}
		command.callback()		
	}

}

type cliCommand struct {
	name string
	description string
	callback func() error
}


func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name: "help",
			description: "Displays the available commands",
			callback: callbackHelp,
		},
		"exit": {
			name: "exit",
			description: "Exits the program",
			callback: callbackExit,
		},
	}
}

//Cleans up user input and returns each word as a Token
func cleanInput(str string) []string {
	loweredStr := strings.ToLower(str)
	words := strings.Fields(loweredStr)
	return words
}