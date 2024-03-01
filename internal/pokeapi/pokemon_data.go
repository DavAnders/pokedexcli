package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PokemonData struct {
	Name           string `json:"name"`
	BaseExperience string `json:"base_experience"`
}

type UserData struct {
	Pokedex map[string]PokemonData
}

func (c *Client) FetchPokemonByName(name string) (PokemonData, error) {
	var pokemonData PokemonData
	url := fmt.Sprintf("%s/pokemon/%s", baseURL, name)

	resp, err := http.Get(url)
	if err != nil {
		return pokemonData, fmt.Errorf("error fetching Pokemon data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return pokemonData, fmt.Errorf("error with PokeAPI status code: %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&pokemonData)
	if err != nil {
		return pokemonData, fmt.Errorf("error decoding Pokemon data: %v", err)
	}
	return pokemonData, nil
}
