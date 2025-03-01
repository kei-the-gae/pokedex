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

		// get command
		commandName := words[0]

		fmt.Printf("Your command was: %s\n", commandName)
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
