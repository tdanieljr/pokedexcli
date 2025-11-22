package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/tdanieljr/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}
type config struct {
	Next     string
	Previous string
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
}

func commandMap(c *config) error {
	results, err := pokeapi.GetAreas(c.Next)
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
func commandMapb(c *config) error {
	results, err := pokeapi.GetAreas(c.Previous)
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
func commandExit(c *config) error {
	fmt.Println("\nClosing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
func commandHelp(*config) error {
	fmt.Printf(
		`
Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
`)
	return nil
}
