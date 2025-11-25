package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/tdanieljr/pokedexcli/internal/pokecache"
)

type Client struct {
	Cache  *pokecache.Cache
	Client *http.Client
}

type PokeAPI struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func NewClient(interval time.Duration) Client {
	c := pokecache.NewCache(interval)
	cl := http.DefaultClient
	return Client{Cache: c, Client: cl}
}

func (c *Client) GetAreas(url string) (PokeAPI, error) {
	var response *http.Response
	var err error
	body, ok := c.Cache.Get(url)
	if ok {
		var results PokeAPI
		json.Unmarshal(body, &results)
		return results, nil
	}

	if url == "" {
		response, err = c.Client.Get("https://pokeapi.co/api/v2/location-area?limit=20")
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
	body, err = io.ReadAll(response.Body)
	if err != nil {
		return PokeAPI{}, err
	}
	c.Cache.Add(url, body)
	var results PokeAPI
	json.Unmarshal(body, &results)
	return results, nil
}
