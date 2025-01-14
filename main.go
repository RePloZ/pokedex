package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/mtslzr/pokeapi-go"
)

type cliCommand struct {
	name        string
	description string
	callback    func(args []string) error
}

var mapIndex int = 0

var catchedPokemon map[string]string = make(map[string]string, 0)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanText := cleanInput(input)
		commands := getCommands()
		if command, exists := commands[cleanText[0]]; exists {
			command.callback(cleanText[1:])
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
		"explore": {
			name:        "explore",
			description: "Explore pokemon in a given location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "cath the given name pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Show the information of a given pokemon name",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Show a list of all the names of the captured pokemon",
			callback:    commandPokedex,
		},
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func commandHelp(args []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	commands := getCommands()
	for _, command := range commands {
		fmt.Printf("%v: %v \n", command.name, command.description)
	}
	return nil
}

func commandExit(args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap(args []string) error {
	locationArea, err := pokeapi.Resource("location-area", mapIndex*20)
	if err != nil {
		return err
	}

	for _, result := range locationArea.Results {
		fmt.Printf("%v \n", result.Name)
	}

	mapIndex++
	return nil
}

func commandMapB(args []string) error {
	mapIndex--
	locationArea, err := pokeapi.Resource("location-area", mapIndex*20)
	if err != nil {
		return err
	}

	for _, result := range locationArea.Results {
		fmt.Printf("%v \n", result.Name)
	}

	return err
}

func commandExplore(args []string) error {
	locationArea, err := pokeapi.LocationArea(args[0])
	if err != nil {
		return err
	}

	for _, pokemon := range locationArea.PokemonEncounters {
		fmt.Printf("%v \n", pokemon.Pokemon.Name)
	}
	return nil
}

func commandCatch(args []string) error {
	pokemonName := args[0]
	pokemon, err := pokeapi.Pokemon(pokemonName)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %v...\n", pokemon.Name)
	if pokemon.BaseExperience > rand.Intn(100) {
		fmt.Printf("%v escaped!\n", pokemon.Name)
		return nil
	}
	catchedPokemon[pokemon.Name] = pokemon.Name
	fmt.Printf("%v was caught!\n", pokemon.Name)
	return nil
}

func commandInspect(args []string) error {
	pokemonName := args[0]
	if _, exists := catchedPokemon[pokemonName]; !exists {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	pokemon, err := pokeapi.Pokemon(pokemonName)
	if err != nil {
		return err
	}
	//pokemonTypes, err
	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats: ")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types: ")
	for _, types := range pokemon.Types {
		fmt.Printf("  - %v\n", types.Type.Name)
	}
	return nil
}

func commandPokedex(args []string) error {
	fmt.Println("Your Pokedex: ")
	for _, pokemon := range catchedPokemon {
		fmt.Printf("  - %v\n", pokemon)
	}
	return nil
}
