package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strings"

	"github.com/tdanieljr/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}
type config struct {
	Client   pokeapi.Client
	Next     string
	Previous string
	Pokedex  map[string]pokeapi.Pokemon
}

var supportedCommands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"help": {
		name:        "help",
		description: "Print the help text",
		callback:    commandHelp,
	},
	"map": {
		name:        "map",
		description: "Print locations",
		callback:    commandMap,
	},
	"mapb": {
		name:        "mapb",
		description: "Print previous locations",
		callback:    commandMapb,
	},
	"explore": {
		name:        "explore",
		description: "Print pokemon in a location",
		callback:    commandExplore,
	},
	"catch": {
		name:        "catch",
		description: "Catch named pokemon",
		callback:    commandCatch,
	},
	"inspect": {
		name:        "inspect",
		description: "Inspect named pokemon",
		callback:    commandInspect,
	},
	"pokedex": {
		name:        "pokedex",
		description: "Print Pokedex",
		callback:    commandPokedex,
	},
}

func commandInspect(c *config, args ...string) error {
	data, ok := c.Pokedex[args[0]]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Printf("%v\n", data)
	return nil
}
func commandPokedex(c *config, args ...string) error {
	for key := range c.Pokedex {
		fmt.Printf("- %s\n", key)
	}
	return nil
}

func commandCatch(c *config, args ...string) error {
	fmt.Printf("\nThrowing a Pokeball at %s...\n", args[0])
	results, err := c.Client.GetPokemon(args[0])
	if err != nil {
		return err
	}
	baseExp := results.BaseExperience
	throw := rand.IntN(baseExp)

	if throw > baseExp/2 {
		fmt.Printf("%s was caught!\n", args[0])

		pokemon := pokeapi.Pokemon{
			Name:   args[0],
			ID:     results.ID,
			BaseXP: baseExp,
		}
		c.Pokedex[args[0]] = pokemon

	} else {
		fmt.Printf("%s escaped!\n", args[0])

	}

	return nil
}

func commandExplore(c *config, args ...string) error {
	fmt.Printf("Exploring %s...\n", args[0])
	results, err := c.Client.GetLocation(args[0])
	if err != nil {
		return err
	}
	for _, encounter := range results.PokemonEncounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}
	return nil
}
func commandMap(c *config, args ...string) error {
	results, err := c.Client.GetAreas(c.Next)
	if err != nil {
		return err
	}
	c.Next = results.Next
	if results.Previous != nil {
		c.Previous = *results.Previous
	} else {
		c.Previous = ""
	}
	for _, r := range results.Results {
		fmt.Println(r.Name)
	}
	return nil
}
func commandMapb(c *config, args ...string) error {
	results, err := c.Client.GetAreas(c.Previous)
	if err != nil {
		return err
	}
	c.Next = results.Next
	if results.Previous != nil {
		c.Previous = *results.Previous
	} else {
		c.Previous = ""
	}
	for _, r := range results.Results {
		fmt.Println(r.Name)
	}
	return nil
}

func cleanInput(text string) []string {
	text = strings.ToLower(strings.TrimSpace(text))
	s := strings.Fields(text)

	return s
}
func commandExit(c *config, args ...string) error {
	fmt.Println("\nClosing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
func commandHelp(c *config, args ...string) error {
	fmt.Printf(
		`
Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
`)
	return nil
}
