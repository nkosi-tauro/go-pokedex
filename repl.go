package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

)

// StartRepl starts the REPL (Read-Eval-Print-Loop)
func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter text: ")
		scanner.Scan()

		text := scanner.Text()

		cleaned := cleanInput(text)
		var queryStr string;
		// If nothing was entered, continue the loop
		if len(cleaned) == 0 {
			continue
		}

		// If there is more than one word, the second word is the query string
		if len(cleaned) > 1 {
			queryStr = cleaned[1]
		}
		//The first inputed word is the command
		commandName := cleaned[0]
		
		availableCommands := getCommands()
		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("Invalid Command")
			continue
		}
		err := command.callback(cfg, queryStr)		
		if err != nil {
			fmt.Println(err)
		}
	}

}

type cliCommand struct {
	name string
	description string
	callback func(*config, string) error
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
		"map": {
			name: "map",
			description: "Lists all location areas",
			callback: callbackMap,
		},
		"mapb": {
			name: "mapb",
			description: "Lists all previous location areas",
			callback: callbackMapBack,
		},
		"search": {
			name: "search",
			description: "Lists the details of the pokemon with the given name",
			callback: callbackPokeSearch,
		},
	}
}

//Cleans up user input and returns each word as a Token
func cleanInput(str string) []string {
	loweredStr := strings.ToLower(str)
	words := strings.Fields(loweredStr)
	return words
}