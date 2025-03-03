package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// GetPokemon -
func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	// get endpoint url
	url := baseURL + "/pokemon/" + pokemonName

	// get from cache if exists
	if val, ok := c.cache.Get(url); ok {
		pokemonResp := Pokemon{}
		if err := json.Unmarshal(val, &pokemonResp); err != nil {
			return Pokemon{}, err
		}
		return pokemonResp, nil
	}

	// new request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	// set request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	// read response
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	// unmarshal response
	pokemonResp := Pokemon{}
	if err := json.Unmarshal(data, &pokemonResp); err != nil {
		return Pokemon{}, err
	}

	// add to cache and return
	c.cache.Add(url, data)

	return pokemonResp, nil
}
