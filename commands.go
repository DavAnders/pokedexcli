package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/DavAnders/pokedexcli/internal/pokeapi"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*config) error
}

var CommandMap = make(map[string]CliCommand)

func InitializeCommands() {
	CommandMap["help"] = CliCommand{
		Name:        "help",
		Description: "Displays a help message",
		Callback:    Help,
	}
	CommandMap["exit"] = CliCommand{
		Name:        "exit",
		Description: "Exit the Pokedex",
		Callback:    Exit,
	}
	CommandMap["map"] = CliCommand{
		Name:        "map",
		Description: "Lists 20 location areas",
		Callback:    callbackMap,
	}
	CommandMap["mapb"] = CliCommand{
		Name:        "mapb",
		Description: "Lists 20 previous location areas",
		Callback:    mapbFunc,
	}
}

func GetCommand(name string) (CliCommand, bool) {
	cmd, found := CommandMap[name]
	return cmd, found
}

func Exit(cfg *config) error {
	fmt.Println("Exiting...")
	return nil
}

func Help(cfg *config) error {
	fmt.Println("Available commands:")
	for name, cmd := range CommandMap {
		fmt.Printf("%s - %s\n", name, cmd.Description)
	}
	return nil
}

func callbackMap(cfg *config) error {
	var cacheKey string
	var locationAreaResp pokeapi.LocationAreaResp // declare a variable to hold the response

	// determine the cache key
	if cfg.nextLocationAreaURL != nil && *cfg.nextLocationAreaURL != "" {
		cacheKey = *cfg.nextLocationAreaURL
	} else {
		cacheKey = "default_location_area" // use a default key for initial fetch
	}

	// attempt to retrieve data from the cache
	cachedData, found := cfg.pokeapiCache.Get(cacheKey)
	if found {
		fmt.Println("Using cached data for location areas.")
		err := json.Unmarshal(cachedData, &locationAreaResp)
		if err != nil {
			return fmt.Errorf("error unmarshaling cached data: %v", err)
		}
	} else {
		fmt.Println("Fetching data from the PokeAPI.")
		var err error
		locationAreaResp, err = cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL) // fetch and directly store the API response
		if err != nil {
			return err
		}

		// cache new data
		dataToCache, err := json.Marshal(locationAreaResp)
		if err != nil {
			return fmt.Errorf("error marshaling data for cache: %v", err)
		}
		cfg.pokeapiCache.Add(cacheKey, dataToCache)
	}

	printLocationAreas(locationAreaResp)

	updateConfigURLs(cfg, &locationAreaResp)

	return nil
}

func mapbFunc(cfg *config) error {
	var cacheKey string
	var locationAreaResp pokeapi.LocationAreaResp

	// determine the cache key
	if cfg.previousLocationAreaURL != nil && *cfg.previousLocationAreaURL != "" {
		cacheKey = *cfg.previousLocationAreaURL
	} else {
		return errors.New("No previous URL to fetch location areas from")
	}

	cachedData, found := cfg.pokeapiCache.Get(cacheKey)
	if found {
		fmt.Println("Using cached data for location areas.")
		err := json.Unmarshal(cachedData, &locationAreaResp) // Unmarshal into locationAreaResp
		if err != nil {
			return fmt.Errorf("error unmarshaling data: %v", err)
		}
	} else {
		fmt.Println("Fetching previous data from the PokeAPI.")
		var err error
		locationAreaResp, err = cfg.pokeapiClient.ListLocationAreas(&cacheKey) // directly store API response
		if err != nil {
			return err
		}

		dataToCache, err := json.Marshal(locationAreaResp)
		if err != nil {
			return fmt.Errorf("error marshaling data for cache: %v", err)
		}
		cfg.pokeapiCache.Add(cacheKey, dataToCache)
	}

	printLocationAreas(locationAreaResp)
	updateConfigURLs(cfg, &locationAreaResp)

	return nil
}

func printLocationAreas(resp pokeapi.LocationAreaResp) {
	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
}

func updateConfigURLs(cfg *config, resp *pokeapi.LocationAreaResp) {
	cfg.nextLocationAreaURL = resp.Next
	cfg.previousLocationAreaURL = resp.Previous
}
