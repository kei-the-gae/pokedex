package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	// get url
	url := baseURL + "/location-area"
	// if url is passed in argument, use that
	if pageURL != nil {
		url = *pageURL
	}

	// start get request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	// make request to server
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	// read response
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	// unmarshal response
	locationsResp := RespShallowLocations{}
	if err := json.Unmarshal(data, &locationsResp); err != nil {
		return RespShallowLocations{}, err
	}

	return locationsResp, nil
}
