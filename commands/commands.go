package commands

import "fmt"

type CliCommand struct {
	Name        string
	Description string
	Callback    func() error
}

var CommandMap = make(map[string]CliCommand)

func InitializeCommands() {
	CommandMap["help"] = CliCommand{
		Name:        "help",
		Description: "Displays a help message",
		Callback:    Help,
	}
	CommandMap["exit"] = CliCommand{
		Name:        "exit",
		Description: "Exit the REPL",
		Callback:    Exit,
	}
}

func GetCommand(name string) (CliCommand, bool) {
	cmd, found := CommandMap[name]
	return cmd, found
}

func Exit() error {
	fmt.Println("Exiting...")
	return nil
}

func Help() error {
	fmt.Println("Available commands:")
	for name, cmd := range CommandMap {
		fmt.Printf("%s - %s\n", name, cmd.Description)
	}
	return nil
}
