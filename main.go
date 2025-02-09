package main

import (
	"github.com/grahamchill/pokedex/internal/functions"
	"github.com/grahamchill/pokedex/internal/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &functions.Config{
		PokeapiClient: pokeClient,
	}

	functions.StartRepl(cfg)
}
