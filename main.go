package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mtslzr/pokeapi-go"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var mapIndex int = 0

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanText := cleanInput(input)
		//fmt.Printf("Your command was: %v\n", cleanText[0])

		commands := getCommands()
		for _, command := range commands {
			if command.name == cleanText[0] {
				command.callback()
			}
		}
	}

}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display the name of the 20 next location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the name of the 20 previous location areas in the Pokemon world",
			callback:    commandMapB,
		},
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func commandHelp() error {
	fmt.Println(`Welcome to the Pokedex!
Usage:
`)

	commands := getCommands()
	for _, command := range commands {
		fmt.Printf("%v: %v \n", command.name, command.description)
	}
	return nil
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap() error {
	locationArea, err := pokeapi.Resource("location-area", mapIndex*20)

	for _, result := range locationArea.Results {
		fmt.Printf("%v \n", result.Name)
	}

	mapIndex++
	return err
}

func commandMapB() error {
	mapIndex--
	locationArea, err := pokeapi.Resource("location-area", mapIndex*20)
	for _, result := range locationArea.Results {
		fmt.Printf("%v \n", result.Name)
	}

	return err
}
