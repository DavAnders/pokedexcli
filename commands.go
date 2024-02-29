package main

import (
	"errors"
	"fmt"
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
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}
	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	cfg.nextLocationAreaURL = resp.Next
	cfg.previousLocationAreaURL = resp.Previous
	return nil
}

func mapbFunc(cfg *config) error {
	if cfg.previousLocationAreaURL == nil {
		return errors.New("already at first page of entries for location areas")
	}
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.previousLocationAreaURL)
	if err != nil {
		return err
	}
	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	cfg.nextLocationAreaURL = resp.Next
	cfg.previousLocationAreaURL = resp.Previous
	return nil
}
