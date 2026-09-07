package main

import (
	"fmt"
)

func commandMapf(cfg *config, args ...string) error {
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

func commandMapb(cfg *config, args ...string) error {
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
