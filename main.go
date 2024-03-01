package main

import (
	"time"

	"github.com/DavAnders/pokedexcli/internal/pokeapi"
	"github.com/DavAnders/pokedexcli/internal/pokecache"
)

type config struct {
	pokeapiClient           pokeapi.Client
	nextLocationAreaURL     *string
	previousLocationAreaURL *string
	pokeapiCache            *pokecache.Cache
	Pokedex                 map[string]pokeapi.PokemonData
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
		pokeapiCache:  pokecache.NewCache(time.Minute * 5),
		Pokedex:       make(map[string]pokeapi.PokemonData),
	}
	start(&cfg)
}
