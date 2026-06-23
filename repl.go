package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/invaderjt/pokedex/internal/pokeapi"
)

func startRepl(config *Config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			fmt.Printf("Invalid input: %v\n", scanner.Err())
			continue
		}
		input := cleanInput(scanner.Text())[0]
		commands := getCommands()
		if command, ok := commands[input]; ok {
			command.callback(config)
			continue
		} else {
			fmt.Println("Unknown command")
		}
	}
}

type Config struct {
	pokeapiClient pokeapi.Client
	nextLocation  *string
	prevLocation  *string
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func commandExit(config *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *Config) error {
	fmt.Print("\nWelcome to the Pokedex!\nUsage:\n\n")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandMap(config *Config) error {
	locationsResp, err := config.pokeapiClient.LocationsList(config.nextLocation)
	if err != nil {
		return err
	}

	config.nextLocation = locationsResp.Next
	config.prevLocation = locationsResp.Previous

	for _, location := range locationsResp.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapb(config *Config) error {
	if config.prevLocation == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	config.nextLocation = config.prevLocation
	return commandMap(config)
}
