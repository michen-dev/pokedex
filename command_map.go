package main

import (
	"fmt"

	"github.com/michen-dev/pokedex/internal/pokeAPI"
)

func commandMapf(cfg *config, pokedex map[string]pokeAPI.Poke_Detail, args ...string) error {
	data, err := cfg.pokeapiClient.Get_location_areas(cfg.next)
	if err != nil {
		return err
	}

	cfg.next = data.Next
	cfg.prev = data.Previous

	location_areas := data.Results
	for _, d := range location_areas {
		fmt.Println(d.Name)
	}
	return nil
}

func commandMapb(cfg *config, pokedex map[string]pokeAPI.Poke_Detail, args ...string) error {
	if cfg.prev == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	data, err := cfg.pokeapiClient.Get_location_areas(cfg.prev)
	if err != nil {
		return err
	}

	cfg.next = data.Next
	cfg.prev = data.Previous

	location_areas := data.Results
	for _, d := range location_areas {
		fmt.Println(d.Name)
	}
	return nil
}
