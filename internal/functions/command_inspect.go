package functions

import (
	"errors"
	"fmt"
	"unicode"
)

// CommandInspect - Inspect caught Pokemon
func CommandInspect(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	checkPokemon := args[0]
	if pokemon, exists := cfg.CaughtPokemon[checkPokemon]; exists {
		// unicode.ToUpper(rune(s[0])) + return string(firstLetter) + s[1:]
		name := string(unicode.ToUpper(rune(pokemon.Name[0]))) + pokemon.Name[1:]
		fmt.Printf("Name: \t %s \n", name)
		fmt.Printf("Height: \t %v \n", pokemon.Height)
		fmt.Printf("Weight: \t %v \n", pokemon.Weight)
		fmt.Printf("Stats: \n")
		for _, stat := range pokemon.Stats {
			fmt.Printf(" - %s: %v \n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Printf("Type: \n")
		for _, pokemonType := range pokemon.Types {
			fmt.Printf(" - %s\n", pokemonType.Type.Name)
		}

		return nil
	} else {
		str := fmt.Sprintf("You have not caught %s that Pokemon yet", checkPokemon)
		return errors.New(str)
	}
}
