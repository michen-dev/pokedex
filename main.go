package main

import (
	"fmt"
	"bufio"
	"os"
)


type cliCommand struct {
	name string
	description string
	callback func() error
}


func main() {
	commands := getCommands()
	commands["help"].callback()
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
			registry.callback()
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("reading input error: %v", err)
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand {
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
	}

}