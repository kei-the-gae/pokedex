package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	// check argument length
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	// get pokemon info
	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	// calculate catch based on experience
	res := rand.Intn(pokemon.BaseExperience)

	// throw pokeball
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if res > 40 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)
	fmt.Println("You may now inspect it with the inspect command.")

	// add pokemon to caught list
	cfg.caughtPokemon[pokemon.Name] = pokemon
	return nil
}
