package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		moi := cleanInput(input)
		fmt.Printf("Your command was: %v\n", moi[0])
	}

}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
