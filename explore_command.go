package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

func explore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("no location specified")
	}
	locationAreaName := args[0]
	cacheKey := "explore_" + locationAreaName

	if cacheData, found := cfg.pokeapiCache.Get(cacheKey); found {
		fmt.Println("Using cached data for:", locationAreaName)
		//
	} else {
		fmt.Println("Fetching data for:", locationAreaName)
		resp, err := cfg.pokeapiClient.FetchLocationAreaByName(locationAreaName)
		if err != nil {
			return fmt.Errorf("error fetching location area details: %v", err)
		}
		//
		dataToCache, err := json.Marshal(resp)
		if err != nil {
			return fmt.Errorf("error marshaling data for cache: %v", err)
		}
		cfg.pokeapiCache.Add(cacheKey, dataToCache)
	}
	return nil
}
