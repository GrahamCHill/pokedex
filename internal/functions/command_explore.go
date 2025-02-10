package functions

import (
	"errors"
	"fmt"
)

func CommandExplore(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a valid location name from map or mapb")
	}

	name := args[0]
	location, err := cfg.PokeapiClient.GetLocation(name)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", location.Name)
	fmt.Println("Found the following Pokemon: ")
	for _, enc := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}
	return nil
}
