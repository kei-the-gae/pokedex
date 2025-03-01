package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// GetLocation -
func (c *Client) GetLocation(locationName string) (Location, error) {
	url := baseURL + "/location-area/" + locationName

	// get from cache
	if val, ok := c.cache.Get(url); ok {
		locationResp := Location{}
		if err := json.Unmarshal(val, &locationResp); err != nil {
			return Location{}, err
		}
		return locationResp, nil
	}

	// start get request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	// make request to server
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer resp.Body.Close()

	// read response
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}

	// unmarshal response
	locationResp := Location{}
	if err := json.Unmarshal(data, &locationResp); err != nil {
		return Location{}, err
	}

	// add data to cache and return
	c.cache.Add(url, data)
	return locationResp, nil
}
