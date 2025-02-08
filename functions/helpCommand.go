package functions

import (
	"fmt"
)

func HelpCommand(commands map[string]CliCommand) error {
	fmt.Println("Available commands:")
	for _, cmd := range commands {
		fmt.Printf("  %s - %s\n", cmd.Name, cmd.Description)
	}
	return nil
}
