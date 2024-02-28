package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/DavAnders/pokedexcli/commands"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin) //read input
	fmt.Println("PokeCLI (type 'exit' to quit)")
	for scanner.Scan() {
		input := scanner.Text()
		command, found := commands.GetCommand(input)
		if found {
			err := command.Callback()
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error", err) // continue around here
			}
		}
		fmt.Printf("You entered: %s\n", input)
		fmt.Println("Enter another command:")
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
