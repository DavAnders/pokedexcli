package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/DavAnders/pokedexcli/internal/pokeapi"
)

func explore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("no location specified")
	}
	locationAreaName := args[0]
	cacheKey := "explore_" + locationAreaName

	var detail pokeapi.LocationAreaDetail

	if cacheData, found := cfg.pokeapiCache.Get(cacheKey); found {
		fmt.Println("Using cached data for:", locationAreaName)
		// deserialize cached data
		err := json.Unmarshal(cacheData, &detail)
		if err != nil {
			return fmt.Errorf("error unmarshaling cached data %v", err)
		}
	} else {
		fmt.Println("Fetching data for:", locationAreaName)
		// fetch from api
		var err error
		detail, err = cfg.pokeapiClient.FetchLocationAreaDetail(locationAreaName)
		if err != nil {
			return fmt.Errorf("error fetching location area details: %v", err)
		}
		// marshal fetched data, add to cache
		dataToCache, err := json.Marshal(detail)
		if err != nil {
			return fmt.Errorf("error marshaling data for cache: %v", err)
		}
		cfg.pokeapiCache.Add(cacheKey, dataToCache)
	}

	// print pokemon encounters
	fmt.Println("Exploring", locationAreaName+"...")
	fmt.Println("Found Pokemon:")
	for _, encounter := range detail.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil
}
