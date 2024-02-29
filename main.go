package main

import "github.com/DavAnders/pokedexcli/internal/pokeapi"

type config struct {
	pokeapiClient           pokeapi.Client
	nextLocationAreaURL     *string
	previousLocationAreaURL *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
	}
	start(&cfg)

}
