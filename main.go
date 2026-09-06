package main

import (
	"time"
	"github.com/michen-dev/pokedex/internal/pokeAPI"
)


type cliCommand struct {
	name string
	description string
	callback func(*config) error
}

type config struct {
	commands map[string]cliCommand
	pokeapiClient pokeAPI.Client
	next *string
	prev *string
}


func main() {
	cfg := config{
		commands: getCommands(),
		pokeapiClient: pokeAPI.NewClient(5 * time.Second, time.Second*5),
	}

	startRepl(&cfg)
}
