package pokeapi

import (
	"encoding/json"
	"fmt"
	"github.com/tdanieljr/pokedexcli/internal/pokecache"
	"io"
	"net/http"
	"time"
)

type Client struct {
	Cache  *pokecache.Cache
	Client *http.Client
}
type Pokemon struct {
	Name   string
	ID     int
	BaseXP int
}

func NewClient(interval time.Duration) Client {
	c := pokecache.NewCache(interval)
	cl := http.DefaultClient
	return Client{Cache: c, Client: cl}
}
func (c *Client) GetPokemon(pokemon string) (PokeData, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", pokemon)
	response, err := http.Get(url)
	if err != nil {
		return PokeData{}, err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return PokeData{}, err
	}
	c.Cache.Add(url, body)
	var results PokeData
	json.Unmarshal(body, &results)
	return results, nil
}

func (c *Client) GetLocation(area string) (PokeLocation, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", area)

	body, ok := c.Cache.Get(url)
	if ok {
		var results PokeLocation
		json.Unmarshal(body, &results)
		return results, nil
	}
	response, err := http.Get(url)
	if err != nil {
		return PokeLocation{}, err
	}
	defer response.Body.Close()
	body, err = io.ReadAll(response.Body)
	if err != nil {
		return PokeLocation{}, err
	}
	c.Cache.Add(url, body)
	var results PokeLocation
	json.Unmarshal(body, &results)
	return results, nil
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
