package main

import (
	"bufio"
	"fmt"
	"github.com/grahamchill/pokedex/functions"
	"os"
	"strings"
)

func main() {
	cfg := &functions.Config{
		Commands: functions.Commands, // Set Commands here
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
		command, exists := functions.Commands[commandName] // Access Commands via Config
		if !exists {
			fmt.Printf("Unknown command: %s\n", commandName)
			continue
		}

		// Execute the command, passing Config (which includes Commands)
		if err := command.Callback(cfg); err != nil {
			fmt.Printf("Error executing command '%s': %v\n", commandName, err)
		}
	}
}
