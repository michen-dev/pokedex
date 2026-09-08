package main

import (
	"fmt"

	"github.com/michen-dev/pokedex/internal/pokeAPI"
)

func commandExplore(cfg *config, pokedex map[string]pokeAPI.Poke_Detail, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("Please provide location area")
	}
	location_area := args[0]
	data, err := cfg.pokeapiClient.Get_pokemon_list(location_area)
	if err != nil {
		return err
	}
	
	fmt.Printf("Exploring %s...\n", location_area)
	fmt.Println("Found Pokemon:")
	for _, pokemon := range data.Pokemon_encounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}
	return nil
}