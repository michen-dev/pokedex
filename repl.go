package main 

import (
	"bufio"
	"fmt"
	"os"
)


func startRepl(cfg *config) {
	commands := cfg.commands
	commandHelp(cfg)
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
		if registry, ok := commands[command]; !ok {
			fmt.Println("Unknown command")
		} else {
			registry.callback(cfg)
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
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
	}
}