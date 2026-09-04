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