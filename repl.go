package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/DavAnders/pokedexcli/commands"
)

func start() {
	scanner := bufio.NewScanner(os.Stdin)
	commands.InitializeCommands()
	for {
		fmt.Print("> ")
		scanner.Scan()
		text := scanner.Text()

		parts := strings.Fields(text) // separating to handle more than one argument
		if len(parts) == 0 {
			continue
		}
		command := parts[0]

		if command == "exit" {
			fmt.Println("Exiting...")
			break
		}

		if cmd, found := commands.GetCommand(command); found {
			err := cmd.Callback()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error executing command %v\n", err)
			}
			continue
		} else {
			fmt.Println("Unknown command:", command)
		}
	}

}
