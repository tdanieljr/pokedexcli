package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Pokedex >")
	scanner := bufio.NewScanner(os.Stdin)
	c := &config{}
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
		err := callBack.callback(c)
		if err != nil {
			fmt.Printf("%v\n", err)
			fmt.Print("Pokedex >")
		}
		fmt.Print("Pokedex >")
	}

}
