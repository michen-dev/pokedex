package main

import (
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
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