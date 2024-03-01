package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"

	"github.com/DavAnders/pokedexcli/internal/pokeapi"
)

func catch(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("no Pokemon specified")
	}
	pokemonName := args[0]
	if _, exists := cfg.Pokedex[pokemonName]; exists {
		fmt.Println("You've already caught", pokemonName)
		return nil
	}

	cacheKey := "pokemon_" + pokemonName
	var pokemonData pokeapi.PokemonData
	if cacheData, found := cfg.pokeapiCache.Get(cacheKey); found {
		err := json.Unmarshal(cacheData, &pokemonData)
		if err != nil {
			return fmt.Errorf("error unmarshaling cached data: %v", err)
		}
	} else {
		var err error
		pokemonData, err = cfg.pokeapiClient.FetchPokemonByName(pokemonName)
		if err != nil {
			return fmt.Errorf("error fetching Pokemon data from API: %v", err)
		}

		dataToCache, err := json.Marshal(pokemonData)
		if err != nil {
			return fmt.Errorf("error marshaling data for cache: %v", err)
		}
		cfg.pokeapiCache.Add(cacheKey, dataToCache)
	}

	catchRate := calculateCatchChance(pokemonData.BaseExperience)
	if rand.Float64() < catchRate {
		// successful catch & add to pokedex
		cfg.Pokedex[pokemonName] = pokemonData
		fmt.Println(cfg.Pokedex[pokemonName].Name, "was caught!")
	} else {
		fmt.Println(pokemonData.Name, "escaped!")
	}
	return nil
}

func calculateCatchChance(baseExp int) float64 {
	var catchRate float64
	switch {
	case baseExp > 200:
		// Hard case
		catchRate = 0.3
	case baseExp >= 150:
		// Medium case
		catchRate = 0.5
	default:
		// Easy
		catchRate = 0.7
	}
	return catchRate
}
