package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	// wait for user input
	reader := bufio.NewScanner(os.Stdin)

	for {
		// prompt
		fmt.Print("Pokedex > ")
		reader.Scan()

		// get input from scanner
		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		// get command and run
		cmdName := words[0]

		if cmd, ok := getCommands()[cmdName]; ok {
			if err := cmd.callback(); err != nil {
				fmt.Println("Error: ", err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	// initialize commands
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
