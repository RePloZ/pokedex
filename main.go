package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print("Hello, World!")
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
