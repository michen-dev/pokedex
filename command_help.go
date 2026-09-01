package main

import (
	"fmt"
)

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	commands := getCommands()
	for command, registry := range commands {
		fmt.Printf("%s: %s\n", command, registry.description)
	}
	return nil
}