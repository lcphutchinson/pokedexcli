package main

import (
	"fmt"
	"os"
	"github.com/lcphutchinson/pokecache"
	"github.com/lcphutchinson/pokecaller"
)

type cliCommand struct {
	name		string
	description	string
	callback	func() error
}

type config struct {

}

func getCommandMap() map[string]cliCommand {
	return map[string]cliCommand {
		"exit": {
			name:		"exit",
			description:	"Exits the session",
			callback:	func() error {
				fmt.Println("Thank you for using PokedexCLI!")
				os.Exit(0)
				return nil
			},
		},
		"help": {
			name:		"help",
			description:	"Displays a help message",
			callback:	func() error {
				fmt.Println("No help subcommand structure yet")
				return nil
			},
		},
		"map": {
			name:		"map",
			description:	"fetches a new page of location areas from PokeAPI",
			callback:	func() error {
				callForMaps()
				return nil
			},
		},
	}
}


