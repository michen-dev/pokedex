package main

import (
	"fmt"

	"github.com/michen-dev/pokedex/internal/pokeAPI"
)

func commandHelp(cfg *config, pokedex map[string]pokeAPI.Poke_Detail, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for _, cmd := range cfg.commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}