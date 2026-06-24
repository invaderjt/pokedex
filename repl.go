package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
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
		input := cleanInput(scanner.Text())
		args := []string{}
		if len(input) > 1 {
			args = input[1:]
		}
		commands := getCommands()
		if command, ok := commands[input[0]]; ok {
			command.callback(config, args...)
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
	myPokedex     map[string]pokeapi.Pokemon
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func commandExit(config *Config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *Config, args ...string) error {
	fmt.Print("\nWelcome to the Pokedex!\nUsage:\n\n")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandMap(config *Config, args ...string) error {
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

func commandMapb(config *Config, args ...string) error {
	if config.prevLocation == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	config.nextLocation = config.prevLocation
	return commandMap(config)
}

func commandExplore(config *Config, args ...string) error {
	if len(args) < 1 {
		fmt.Println("Please provide an area name")
		return errors.New("no area name")
	}
	targetArea := args[0]
	fmt.Printf("Exploring %s...\n", targetArea)

	exploreResp, err := config.pokeapiClient.ExploreLocation(targetArea)
	if err != nil {
		fmt.Printf("%s area not found\n", targetArea)
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, pokemon := range exploreResp.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}
	return nil
}

func commandCatch(config *Config, args ...string) error {
	if len(args) < 1 {
		fmt.Println("Please provide a pokemon name")
		return errors.New("no pokemon name")
	}
	target := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", target)
	pokemonResp, err := config.pokeapiClient.CheckPokemon(target)
	if err != nil {
		fmt.Printf("No such pokemon called %s\n", target)
		return err
	}

	xp := pokemonResp.BaseExperience
	catchRate := (xp / 4) + 50
	caught := false
	if rand.Intn(xp) < catchRate {
		caught = true
	}
	if !caught {
		fmt.Printf("%s escaped!\n", target)
		return nil
	}

	fmt.Printf("%s was caught!\n", target)
	config.myPokedex[target] = pokemonResp
	fmt.Printf("%s added to PokeDex\n", target)
	return nil

}

func commandInspect(config *Config, args ...string) error {
	if len(args) < 1 {
		fmt.Println("Please provide a pokemon name")
		return errors.New("no pokemon name")
	}

	name := args[0]
	if _, ok := config.myPokedex[name]; !ok {
		fmt.Printf("You have not caught a %s\n", name)
		return nil
	}

	pokemon := config.myPokedex[name]
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, value := range pokemon.Stats {
		fmt.Printf(" - %v: %v\n", value.Stat.Name, value.BaseStat)
	}
	fmt.Println("Type(s):")
	for _, pkType := range pokemon.Types {
		fmt.Printf(" - %v\n", pkType.Type.Name)
	}
	return nil

}

func commandPokedex(config *Config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range config.myPokedex {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}
