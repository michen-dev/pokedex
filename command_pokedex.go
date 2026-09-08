package main

import (
	"github.com/michen-dev/pokedex/internal/pokeAPI"
	"fmt"
)

func commandPokedex(cfg *config, pokedex map[string]pokeAPI.Poke_Detail, args ...string) error {
	fmt.Println("Your Pokemon:")
	for name := range pokedex {
		fmt.Printf("- %s\n", name)
	}
	return nil
}