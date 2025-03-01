package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/kei-the-gae/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *config) {
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
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		if cmd, ok := getCommands()[cmdName]; ok {
			if err := cmd.callback(cfg, args...); err != nil {
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
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	// initialize commands
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Explore a location",
			callback:    commandExplore,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
