package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type PokeAPI struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetAreas(url string) (PokeAPI, error) {
	var response *http.Response
	var err error
	if url == "" {
		response, err = http.Get("https://pokeapi.co/api/v2/location-area?limit=20")
		if err != nil {
			return PokeAPI{}, err
		}
	} else {

		response, err = http.Get(url)
		if err != nil {
			return PokeAPI{}, err
		}
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return PokeAPI{}, err
	}
	var results PokeAPI
	json.Unmarshal(body, &results)
	return results, nil
}
