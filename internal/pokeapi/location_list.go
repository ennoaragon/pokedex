package pokeapi

import (
	"encoding/json"
	"net/http"
	"io"
)


func (c *Client) ListLocations(pageUrl *string) (PokemonLocationResponse, error) {
	url := baseURL + "/location-area"

	if pageUrl != nil {
		url = *pageUrl
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonLocationResponse{}, err
	}


	res, err := c.httpClient.Do(req)

	if err != nil {
		return PokemonLocationResponse{}, err
	}

	defer res.Body.Close()

	var body PokemonLocationResponse

	jsonData, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonLocationResponse{}, err
	}

	err = json.Unmarshal(jsonData, &body)

	if err != nil {
		return PokemonLocationResponse{}, err
	}

	return body, nil
}
