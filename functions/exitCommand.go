package functions

import (
	"fmt"
	"os"
)

func ExitCommand(commands map[string]CliCommand) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
