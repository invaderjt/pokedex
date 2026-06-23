package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type Config struct {
	Next_Location *string
	Prev_Location *string
}

var first_location = "https://pokeapi.co/api/v2/location-area/"

var config = Config{
	Next_Location: &first_location,
	Prev_Location: nil,
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
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandMap(config *Config) error {
	resp, err := http.Get(*config.Next_Location)
	if err != nil {
		return err
	}

	var locations Locations
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&locations); err != nil {
		return err
	}

	config.Next_Location = &locations.Next
	config.Prev_Location = &locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapb(config *Config) error {
	if *config.Prev_Location == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	config.Next_Location = config.Prev_Location
	return commandMap(config)

}
