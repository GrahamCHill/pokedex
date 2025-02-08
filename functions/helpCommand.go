package functions

import (
	"fmt"
)

func HelpCommand(cfg *Config) error {
	commands := cfg.Commands // Access Commands via Config
	fmt.Println("Available commands:")
	for _, cmd := range commands {
		fmt.Printf("  %s - %s\n", cmd.Name, cmd.Description)
	}
	return nil
}
