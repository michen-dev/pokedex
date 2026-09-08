package main

import (
	"fmt"
	"math/rand"

	"github.com/michen-dev/pokedex/internal/pokeAPI"
)

const (
	catch_threshold = 50
)

func commandCatch(cfg *config, pokedex map[string]pokeAPI.Poke_Detail, args ...string) error {	
	if len(args) == 0 {
		return fmt.Errorf("Please provide Pokemon name")
	}

	name := args[0]
	if _, ok := pokedex[name]; ok {
		fmt.Printf("You already caught %s\n", name)
		return nil
	}

	pokemon, err := cfg.pokeapiClient.Get_pokemon_detail(name)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	base_experience := pokemon.Base_experience
	threshold := catch_threshold + catch_threshold*(base_experience/100)
	roll := rand.Float64() * 100.0
	fmt.Printf("threshold: %v - roll: %v - exp: %v\n", threshold, roll, base_experience)
	if roll >= float64(threshold) {
		fmt.Printf("%s was caught!\n", name)
		pokedex[name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", name)
	}
	return nil
}