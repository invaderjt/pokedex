package main

import (
	"time"

	"github.com/invaderjt/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &Config{
		pokeapiClient: pokeClient,
		myPokedex:     map[string]pokeapi.Pokemon{},
	}

	startRepl(cfg)
}
