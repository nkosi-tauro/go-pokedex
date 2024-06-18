package main

import (
	"bufio"
	"fmt"
	"os"
)

// StartRepl starts the REPL (Read-Eval-Print-Loop)
func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter text: ")
		scanner.Scan()
		text := scanner.Text()
		fmt.Println("echoing", text)
	}

}