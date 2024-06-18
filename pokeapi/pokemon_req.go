package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemonByName(name string) (PokemonResp, error) {
	endpoint := "/pokemon/" + name
	fullURL := baseURL + endpoint

	// make request
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return PokemonResp{}, err
	}

	// get response and close connection
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return PokemonResp{}, fmt.Errorf("The pokemon named %s was not found.", name)
		}
		return PokemonResp{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonResp{}, err
	}

	pokemonResp := PokemonResp{}
	err = json.Unmarshal(data, &pokemonResp)
	if err != nil {
		return PokemonResp{}, err
	}

	return pokemonResp, nil

}