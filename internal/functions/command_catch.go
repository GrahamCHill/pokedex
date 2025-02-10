package functions

import (
	"errors"
	"fmt"
	"math/rand"
)

func CommandCatch(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	pokemon, err := cfg.PokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	res := rand.Intn(pokemon.BaseExperience)

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if res > 40 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)

	fmt.Printf("You may inspect it with the inspect command,\n" +
		"You may also check your pokedex to see all caught pokemon\n")

	cfg.CaughtPokemon[pokemon.Name] = pokemon
	return nil
}
