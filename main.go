package main

import (
	"github.com/grahamchill/pokedex/internal/functions"
	"github.com/grahamchill/pokedex/internal/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &functions.Config{
		CaughtPokemon: map[string]pokeapi.Pokemon{},
		PokeapiClient: pokeClient,
	}

	functions.StartRepl(cfg)
}
