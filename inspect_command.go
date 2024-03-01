package main

import (
	"fmt"
	"strings"
)

func inspectPokemon(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("no Pokemon specified")
	}
	pokemonName := strings.ToLower(args[0])

	pokemonData, exists := cfg.Pokedex[pokemonName]
	if !exists {
		fmt.Printf("You have not caught %s yet.\n", args[0])
		return nil
	}
	fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\nStats:\n", pokemonData.Name, pokemonData.Height, pokemonData.Weight)
	for _, stat := range pokemonData.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokemonType := range pokemonData.Types {
		fmt.Printf("  - %s\n", pokemonType.Type.Name)
	}
	return nil
}
