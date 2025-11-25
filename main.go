package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/tdanieljr/pokedexcli/internal/pokeapi"
)

func main() {
	client := pokeapi.NewClient(5 * time.Second)
	pokedex := make(map[string]pokeapi.Pokemon)
	fmt.Print("Pokedex >")
	scanner := bufio.NewScanner(os.Stdin)
	c := &config{Client: client, Pokedex: pokedex}
	for {
		ok := scanner.Scan()
		if !ok {
			return
		}
		s := cleanInput(scanner.Text())
		if len(s) == 0 {
			fmt.Print("Pokedex >")
			continue
		}
		command := s[0]
		callBack, ok := supportedCommands[command]
		if !ok {
			fmt.Println("Unknown command")
			fmt.Print("Pokedex >")
			continue
		}
		if len(s) > 1 {
			err := callBack.callback(c, s[1])
			if err != nil {
				fmt.Printf("%v\n", err)
				fmt.Print("Pokedex >")
			}
		} else {
			err := callBack.callback(c)

			if err != nil {
				fmt.Printf("%v\n", err)
				fmt.Print("Pokedex >")
				continue
			}
		}
		fmt.Print("Pokedex >")
	}

}
