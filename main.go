package main

import (
	"bufio"
	"fmt"
	"github.com/grahamchill/pokedex/functions"
	"os"
	"strings"
)

func main() {
	commands := map[string]functions.CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    functions.ExitCommand,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    functions.HelpCommand,
		},
	}
	fmt.Printf("Welcome to the Pokedex!\n")
	ioReader := bufio.NewReader(os.Stdin)
	buffer := bufio.NewScanner(ioReader)
	for {
		fmt.Printf("Pokedex > ")
		buffer.Scan()
		text := buffer.Text()
		text = strings.ToLower(text)
		textTrimmed := functions.CleanInput(text)

		commandName := textTrimmed[0]

		// Check if the command exists
		command, exists := commands[commandName]
		if !exists {
			fmt.Printf("Unknown command: %s\n", commandName)
			continue
		}

		// Execute the command
		if err := command.Callback(commands); err != nil {
			fmt.Printf("Error executing command '%s': %v\n", commandName, err)
		}
	}
}
