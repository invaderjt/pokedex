package main

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "List next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Find pokemon in specified area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a specified pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Examine a pokemon you've caught",
			callback:    commandInspect,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}
