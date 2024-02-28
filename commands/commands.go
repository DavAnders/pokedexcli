package commands

type CliCommand struct {
	Name        string
	Description string
	Callback    func() error
}

var commandMap = map[string]CliCommand{
	"help": {
		Name:        "help",
		Description: "Displays a help message",
		Callback:    Help,
	},
	"exit": {
		Name:        "exit",
		Description: "Exit the REPL",
		Callback:    Exit,
	},
}

func GetCommand(name string) (CliCommand, bool) {
	cmd, found := commandMap[name]
	return cmd, found
}
