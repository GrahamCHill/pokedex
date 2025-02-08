package functions

type CliCommand struct {
	Name        string
	Description string
	Callback    func(map[string]CliCommand) error // Update callback signature
}
