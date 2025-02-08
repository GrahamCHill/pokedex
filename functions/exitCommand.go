package functions

import (
	"fmt"
	"os"
)

func ExitCommand(cfg *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
