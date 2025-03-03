package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	// check argument length
	if len(args) != 0 {
		return errors.New("this command does not take any arguments")
	}

	// check if there are any caught pokemon
	if len(cfg.caughtPokemon) == 0 {
		return errors.New("you have not caught any pokemon yet")
	}

	// return pokemon
	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Println("  -", pokemon.Name)
	}
	return nil
}
