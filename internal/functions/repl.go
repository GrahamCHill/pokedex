package functions

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/grahamchill/pokedex/internal/pokeapi"
)

type Config struct {
	PokeapiClient    pokeapi.Client
	NextLocationsURL *string
	PrevLocationsURL *string
	CaughtPokemon    map[string]pokeapi.Pokemon
}

func StartRepl(cfg *Config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := CleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		var args []string
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func CleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    CommandHelp,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Explore a specific location and display the Pokemon",
			callback:    CommandExplore,
		},
		"inspect": {
			name:        "Inspect <pokemon_name>",
			description: "Inspects a specific pokemon, and provides more info on it",
			callback:    CommandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays all of your pokemon",
			callback:    CommandPokedex,
		},
		"catch": {
			name:        "catch <pokemon>",
			description: "Try and Catch a specific pokemon",
			callback:    CommandCatch,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    CommandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    CommandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    CommandExit,
		},
	}
}
