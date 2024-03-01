package main

import (
	"fmt"
	"sort"
)

func listPokedex(cfg *config, args ...string) error {
	if len(cfg.Pokedex) == 0 {
		fmt.Println("Your Pokedex contains no Pokemon. Try using the 'catch' command.")
		return nil
	}

	fmt.Println("Your Pokedex:")
	var pokemonNames []string
	for name := range cfg.Pokedex {
		pokemonNames = append(pokemonNames, name)
	}
	sort.Strings(pokemonNames)

	for _, name := range pokemonNames {
		fmt.Printf(" - %s\n", name)
	}
	return nil
}
