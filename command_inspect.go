package main

import (
	"github.com/michen-dev/pokedex/internal/pokeAPI"
	"fmt"
	"reflect"
)

func commandInspect(cfg *config, pokedex map[string]pokeAPI.Poke_Detail, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("Please provide Pokemon name")
	}

	name := args[0]
	pokemon, ok := pokedex[name]
	if !ok {
		return fmt.Errorf("you have not caught that pokemon")
	}

	v := reflect.ValueOf(pokemon)
	t := reflect.TypeOf(pokemon)

	for i:=0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		fmt.Printf("%s: %v\n", field.Name, value.Interface())
	}

	return nil
}
