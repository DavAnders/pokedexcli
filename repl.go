package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func start(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	InitializeCommands()
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

		if cmd, found := GetCommand(command); found {
			err := cmd.Callback(cfg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error executing command -- %v\n", err)
			}
			continue
		} else {
			fmt.Println("Unknown command:", command)
		}
	}

}
