package functions

// CliCommand represents a command in the CLI.
type CliCommand struct {
	Name        string
	Description string
	Callback    func(*Config) error // Accepts a pointer to Config
}

// Config holds configuration data for the application.
type Config struct {
	NextURL     string
	PreviousURL string
	Commands    map[string]CliCommand
}
