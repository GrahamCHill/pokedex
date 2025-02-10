package functions

import (
	"fmt"
	"unicode"
)

// CommandPokedex - Display all caught pokemon
func CommandPokedex(cfg *Config, args ...string) error {
	fmt.Printf("Pokedex \n")
	// if pokemon, exists := cfg.CaughtPokemon
	for pokemon, _ := range cfg.CaughtPokemon {
		name := string(unicode.ToUpper(rune(pokemon[0]))) + pokemon[1:]
		fmt.Printf(" - %s \n", name)
	}
	return nil
}
