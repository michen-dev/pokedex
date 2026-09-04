package main




type cliCommand struct {
	name string
	description string
	callback func(*config) error
}

type config struct {
	commands map[string]cliCommand
}


func main() {
	cfg := config{
		commands: map[string]cliCommand{
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
		},
	}

	startRepl(&cfg)
}
