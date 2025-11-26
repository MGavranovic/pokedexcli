package main

func main() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		}, "help": {
			name:        "help",
			description: "Displays helpful information about the Pokedex",
			callback:    commandHelp,
		},
	}

	repl()
}
