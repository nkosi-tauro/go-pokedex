package main

import "fmt"

func callbackHelp(cfg *config) error {
	fmt.Println("Welcome to Nkosi's Pokedex.")
	fmt.Println("Here are the available commands: ")
	availableCommands := getCommands()
	for _, cmd := range availableCommands {
		fmt.Printf(" - %s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println("")
	return nil
}