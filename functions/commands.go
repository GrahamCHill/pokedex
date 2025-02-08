package functions

var Commands = map[string]CliCommand{
	"exit": {
		Name:        "exit",
		Description: "Exit the Pokedex",
		Callback:    ExitCommand,
	},
	"help": {
		Name:        "help",
		Description: "Display help information",
		Callback:    HelpCommand,
	},
	"map": {
		Name:        "map",
		Description: "Display 20 Towns and Locations in Pokemon",
		Callback:    MapCommand,
	},
	"mapb": {
		Name:        "mapb",
		Description: "Display previous 20 Towns and Locations in Pokemon",
		Callback:    MapbCommand,
	},
}
