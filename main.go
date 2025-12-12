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
		}, "map": {
			name:        "map",
			description: "Display next 20 locations on the map",
			callback:    commandMap,
		}, "mapb": {
			name:        "mapb",
			description: "Display previous 20 locations on the map",
			callback:    commandMapBack,
		}, "explore": {
			name:        "explore",
			description: "Displays details on an area",
			callback:    commandExplore,
		},
	}

	repl()
}
