package main


import (
	"os"
	"fmt"

	"github.com/michen-dev/pokedex/internal/pokeAPI"
)

func commandExit(cfg *config, pokedex map[string]pokeAPI.Poke_Detail, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}