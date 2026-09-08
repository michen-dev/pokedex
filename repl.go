package main 

import (
	"bufio"
	"fmt"
	"os"
	"github.com/michen-dev/pokedex/internal/pokeAPI"
)


func startRepl(cfg *config) {
	pokedex := map[string]pokeAPI.Poke_Detail{}
	commands := cfg.commands
	commandHelp(cfg, pokedex)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		
		line := scanner.Text()
		words := cleanInput(line)
		if len(words) == 0 {
			continue
		}
		command := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}
		if registry, ok := commands[command]; !ok {
			fmt.Println("Unknown command")
		} else {
			err := registry.callback(cfg, pokedex, args...)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("reading input error: %v", err)
	}
}


func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
				name: "help",
				description: "Displays a help message",
				callback: commandHelp,
			},
		"map": {
			name: "map",
			description: "Display NEXT 20 location areas in Pokemon world",
			callback: commandMapf,
		},
		"mapb": {
			name: "mapb",
			description: "Display PREV 20 location areas in Pokemon world",
			callback: commandMapb,
		},
		"explore": {
			name: "explore",
			description: "Explore the Pokemon list of a specific Location area",
			callback: commandExplore,
		},
		"catch": {
			name: "catch",
			description: "Catch Pokemon",
			callback: commandCatch,
		},
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
	}
}